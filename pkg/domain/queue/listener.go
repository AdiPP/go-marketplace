package queue

import (
	"context"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
)

type Listener interface {
	Handle(ctx context.Context, event event.Event) (err error)
}
