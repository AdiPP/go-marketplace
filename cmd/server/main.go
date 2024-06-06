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
	"reflect"
)

func main() {
	// Context
	ctx := context.Background()

	// Repository Adapter
	dummyRepository := repository.NewDummyRepositoryAdapter()

	// Queue Adapter
	memoryQueueAdapter := infraQueue.NewMemoryQueueAdapter()

	// Use Cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(memoryQueueAdapter, dummyRepository, dummyRepository)
	processPaymentUseCase := usecase.NewProcessOrderPaymentUseCase(memoryQueueAdapter)

	// Handlers
	createOrderHandler := controller.NewCreateOrderHandler(createOrderUseCase)

	// Listeners
	processOrderPaymentListener := listener.NewProcessOrderPaymentListener(processPaymentUseCase)
	stockMovementListener := listener.NewProcessStockMovementListener()

	http.HandleFunc("POST /create-order", createOrderHandler.Handle)

	eventHandlers := map[reflect.Type][]domainQueue.Listener{
		reflect.TypeOf(event.OrderCreatedEvent{}): {
			processOrderPaymentListener,
			stockMovementListener,
		},
	}

	for eventType, handlers := range eventHandlers {
		for _, handler := range handlers {
			memoryQueueAdapter.ListenerRegister(eventType, handler)
		}
	}

	err := memoryQueueAdapter.Connect(ctx)

	if err != nil {
		log.Fatalf("Error connect queue %s", err)
	}

	defer memoryQueueAdapter.Disconnect(ctx)

	go func(ctx context.Context, queueName string) {
		err = memoryQueueAdapter.StartConsuming(ctx, queueName)

		if err != nil {
			log.Fatalf("Error starting consumer %s: %s", queueName, err)
		}
	}(ctx, "default")

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
