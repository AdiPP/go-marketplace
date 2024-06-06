package main

import (
	"context"
	"fmt"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	domainQueue "github.com/AdiPP/go-marketplace/pkg/domain/queue"
	infraQueue "github.com/AdiPP/go-marketplace/pkg/infrastructure/queue"
	"github.com/AdiPP/go-marketplace/pkg/infrastructure/repository"
	"github.com/AdiPP/go-marketplace/pkg/interface/controller"
	"github.com/AdiPP/go-marketplace/pkg/interface/listener"
	"github.com/AdiPP/go-marketplace/pkg/usecase"
	"log"
	"net/http"
)

func main() {
	// Context
	ctx := context.Background()

	// Repository Adapter
	repo := repository.NewDummyRepositoryAdapter()

	// Queue Adapter
	q := infraQueue.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")

	// Use Cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(q, repo, repo)
	processPaymentUseCase := usecase.NewProcessOrderPaymentUseCase(q)

	// Handlers
	createOrderHandler := controller.NewCreateOrderHandler(createOrderUseCase)

	// Listeners
	processOrderPaymentListener := listener.NewProcessOrderPaymentListener(processPaymentUseCase)
	stockMovementListener := listener.NewProcessStockMovementListener()

	http.HandleFunc("POST /create-order", createOrderHandler.Handle)

	eventHandlers := map[string][]domainQueue.Listener{
		event.OrderCreatedEvent{}.GetType(): {
			processOrderPaymentListener,
			stockMovementListener,
		},
	}

	for eventType, handlers := range eventHandlers {
		for _, handler := range handlers {
			q.ListenerRegister(eventType, handler)
		}
	}

	err := q.Connect(ctx)

	if err != nil {
		log.Fatalf("Error connect queue %s", err)
	}

	defer q.Disconnect(ctx)

	orderCreatedEventName := event.OrderCreatedEvent{}.GetType()

	go func(ctx context.Context, queueName string) {
		err = q.StartConsuming(ctx, queueName)

		if err != nil {
			log.Fatalf("Error starting consumer %s: %s", queueName, err)
		}
	}(ctx, orderCreatedEventName)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
