package queue

import (
	"context"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
)

type Publisher interface {
	Publish(ctx context.Context, event event.Event) error
}
