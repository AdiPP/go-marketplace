package listener

import (
	"context"
	"errors"
	"fmt"
)

type ProcessStockMovementListener struct {
}

func NewProcessStockMovementListener() *ProcessStockMovementListener {
	return &ProcessStockMovementListener{}
}

func (p ProcessStockMovementListener) Handle(ctx context.Context, event any) (err error) {
	fmt.Println("--- StockMovementListener ---")

	err = errors.New("process stock movement listener failed")

	if err != nil {
		return
	}

	return nil
}
