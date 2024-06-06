package listener

import (
	"context"
	"fmt"

	"github.com/AdiPP/go-marketplace/pkg/domain/event"
)

type ProcessStockMovementListener struct {
}

func NewProcessStockMovementListener() *ProcessStockMovementListener {
	return &ProcessStockMovementListener{}
}

func (p ProcessStockMovementListener) Handle(ctx context.Context, event event.Event) (err error) {
	fmt.Println("--- StockMovementListener ---")

	return nil
}
