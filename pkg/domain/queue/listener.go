package queue

import (
	"context"
)

type Listener interface {
	Handle(ctx context.Context, event any) (err error)
}
