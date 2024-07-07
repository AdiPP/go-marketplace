package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	domainQueue "github.com/AdiPP/go-marketplace/pkg/domain/queue"
	"github.com/AdiPP/go-marketplace/pkg/infrastructure/config"
	infraQueue "github.com/AdiPP/go-marketplace/pkg/infrastructure/queue"
	"github.com/AdiPP/go-marketplace/pkg/infrastructure/repository"
	"github.com/AdiPP/go-marketplace/pkg/interface/controller"
	"github.com/AdiPP/go-marketplace/pkg/interface/listener"
	"github.com/AdiPP/go-marketplace/pkg/usecase"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// Context
	ctx := context.Background()

	// Config
	dbCfg := config.Database{
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseSchema:   os.Getenv("DATABASE_SCHEMA"),
	}

	// Repository Adapter
	postgresRepo := repository.NewPostgresRepositoryAdapter(dbCfg)
	dummyRepo := repository.NewDummyRepositoryAdapter()

	// Queue Adapter
	q := infraQueue.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")

	// Use Cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(q, dummyRepo, postgresRepo)
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
