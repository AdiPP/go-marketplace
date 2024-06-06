package queue

import (
	"context"
	"encoding/json"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/domain/queue"
	"log"
	"reflect"
	"sync"
)

type MemoryQueueAdapter struct {
	listeners map[string][]queue.Listener
}

func NewMemoryQueueAdapter() *MemoryQueueAdapter {
	return &MemoryQueueAdapter{
		listeners: make(map[string][]queue.Listener),
	}
}

func (m *MemoryQueueAdapter) ListenerRegister(eventType string, listener queue.Listener) {
	m.listeners[eventType] = append(m.listeners[eventType], listener)
}

func (m *MemoryQueueAdapter) Connect(ctx context.Context) error {
	log.Println("MemoryQueueAdapter connected")
	return nil
}

func (m *MemoryQueueAdapter) Disconnect(ctx context.Context) error {
	log.Println("MemoryQueueAdapter disconnected")
	return nil
}

func (m *MemoryQueueAdapter) Publish(ctx context.Context, event event.Event) error {
	eventType := reflect.TypeOf(event)
	payloadJson, _ := json.Marshal(event)

	log.Printf("** [Publish] %s: %v ---", eventType, string(payloadJson))

	var wg sync.WaitGroup

	for _, listener := range m.listeners[eventType.Name()] {
		wg.Add(1)

		go func(listener queue.Listener) {
			defer wg.Done()

			_ = listener.Handle(ctx, event)

			// TODO: Handling failed event

		}(listener)
	}

	wg.Wait()

	return nil
}

func (m *MemoryQueueAdapter) StartConsuming(_ context.Context, queueName string) error {
	log.Printf("MemoryQueueAdapter StartConsuming queue %s\n", queueName)
	return nil
}
