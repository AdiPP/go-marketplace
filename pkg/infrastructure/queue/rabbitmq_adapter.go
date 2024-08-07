package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/domain/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	username  string
	password  string
	host      string
	port      string
	uri       string
	conn      *amqp.Connection
	listeners map[string][]queue.Listener
}

type RabbitMQAdapterOption func(*RabbitMQAdapter)

func RabbitMQAdapterWithUsername(username string) RabbitMQAdapterOption {
	return func(a *RabbitMQAdapter) {
		a.username = username
	}
}

func RabbitMQAdapterWithPassword(password string) RabbitMQAdapterOption {
	return func(a *RabbitMQAdapter) {
		a.password = password
	}
}

func RabbitMQAdapterWithHost(host string) RabbitMQAdapterOption {
	return func(a *RabbitMQAdapter) {
		a.host = host
	}
}

func RabbitMQAdapterWithPort(port string) RabbitMQAdapterOption {
	return func(a *RabbitMQAdapter) {
		a.port = port
	}
}

func NewRabbitMQAdapter(opts ...RabbitMQAdapterOption) *RabbitMQAdapter {
	adapter := &RabbitMQAdapter{
		username:  "guest",
		password:  "guest",
		host:      "localhost",
		port:      "5672",
		uri:       "",
		conn:      &amqp.Connection{},
		listeners: map[string][]queue.Listener{},
	}

	for _, opt := range opts {
		opt(adapter)
	}

	adapter.buildUri()

	return adapter
}

func (r *RabbitMQAdapter) buildUri() {
	r.uri = fmt.Sprintf("amqp://%s:%s@%s:%s/", r.username, r.password, r.host, r.port)
}

func (r *RabbitMQAdapter) Publish(ctx context.Context, event event.Event) error {
	eventType := event.GetType()

	ch, err := r.conn.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		eventType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent to queue %s: %s\n", eventType, body)
	return nil

}

func (r *RabbitMQAdapter) ListenerRegister(eventType string, listener queue.Listener) {
	r.listeners[eventType] = append(r.listeners[eventType], listener)
}

func (r *RabbitMQAdapter) Connect(ctx context.Context) error {
	conn, err := amqp.Dial(r.uri)

	if err != nil {
		return err
	}

	r.conn = conn
	return nil
}

func (r *RabbitMQAdapter) Disconnect(ctx context.Context) error {
	return r.conn.Close()
}

func (r *RabbitMQAdapter) StartConsuming(ctx context.Context, queueName string) error {
	ch, err := r.conn.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		return err
	}

	msgs, err := ch.ConsumeWithContext(ctx, q.Name, "", false, false, false, false, nil)

	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received a message on queue %s: %s", queueName, msg.Body)

			hasError := false

			event, err := event.NewEvent(q.Name)

			if err != nil {
				log.Printf("Failed creating event: %s", err.Error())
				continue
			}

			err = json.Unmarshal(msg.Body, &event)

			if err != nil {
				log.Printf("Failed unmarshal json body: %s", err.Error())
				continue
			}

			eventType := event.GetType()

			handlers, found := r.listeners[eventType]

			if !found {
				log.Printf("Handlers with event type %s not found", eventType)
				continue
			}

			for _, handler := range handlers {
				err := handler.Handle(ctx, event)

				if err != nil {
					log.Printf("Error processing message: %s", err)
					hasError = true
					break
				}
			}

			if !hasError {
				msg.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages on queue %s. To exit press CTRL+C", queueName)
	<-ctx.Done()
	return nil
}
