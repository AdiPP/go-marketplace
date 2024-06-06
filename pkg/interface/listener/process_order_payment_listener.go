package listener

import (
	"context"
	"fmt"
	domainEvent "github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/usecase"
)

type ProcessOrderPaymentListener struct {
	processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase
}

func NewProcessOrderPaymentListener(processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase) *ProcessOrderPaymentListener {
	return &ProcessOrderPaymentListener{processOrderPaymentUseCase: processOrderPaymentUseCase}
}

func (l *ProcessOrderPaymentListener) Handle(ctx context.Context, event domainEvent.Event) (err error) {
	fmt.Println("--- ProcessOrderPaymentListener ---")
	orderCreatedEvent := event.(*domainEvent.OrderCreatedEvent)

	var useCaseOrderItems []usecase.OrderItem

	for _, orderItem := range orderCreatedEvent.Items {
		useCaseOrderItems = append(useCaseOrderItems, usecase.OrderItem{
			ProductName: orderItem.ProductName,
			Quantity:    orderItem.Quantity,
			TotalPrice:  orderItem.TotalPrice,
		})
	}

	err = l.processOrderPaymentUseCase.Execute(ctx, usecase.ProcessOrderPaymentDto{
		Order: usecase.Order{
			Id:         orderCreatedEvent.Id,
			Items:      useCaseOrderItems,
			TotalPrice: orderCreatedEvent.TotalPrice,
			Status:     orderCreatedEvent.Status,
		},
	})

	if err != nil {
		return err
	}

	return
}
