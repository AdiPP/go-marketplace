package queue

import (
	"context"
)

type Queue interface {
	Publisher
	ListenerRegister(eventType string, listener Listener)
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	StartConsuming(ctx context.Context, queueName string) error
}
