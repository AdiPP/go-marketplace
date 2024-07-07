package usecase

import (
	"context"
	"time"

	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/domain/queue"
)

type ProcessOrderPaymentUseCase struct {
	publisher queue.Publisher
}

func NewProcessOrderPaymentUseCase(publisher queue.Publisher) *ProcessOrderPaymentUseCase {
	return &ProcessOrderPaymentUseCase{publisher: publisher}
}

func (u *ProcessOrderPaymentUseCase) Execute(ctx context.Context, dto ProcessOrderPaymentDto) error {
	order := entity.RestoreOrderEntity(dto.Id, dto.Status)

	for _, item := range dto.Items {
		orderItem := entity.NewOrderItemEntity(item.ProductName, item.TotalPrice/float64(item.Quantity), item.Quantity)

		order.AddItem(orderItem)
	}

	paymentValue := dto.TotalPrice

	err := order.Pay(paymentValue)

	if err != nil {
		return err
	}

	err = u.publisher.Publish(ctx, event.OrderPaidEvent{
		OrderId:     order.Id(),
		PaidValue:   paymentValue,
		PaymentDate: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}
