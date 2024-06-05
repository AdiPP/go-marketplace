package queue

import (
	"context"
	"reflect"
)

type Queue interface {
	Publisher
	ListenerRegister(eventType reflect.Type, listener Listener)
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	StartConsuming(ctx context.Context, queueName string) error
}
