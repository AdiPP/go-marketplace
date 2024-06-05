package usecase

import (
	"context"
	"fmt"
	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/domain/queue"
)

type CreateOrderUseCase struct {
	publisher queue.Publisher
}

func NewCreateOrderUseCase(publisher queue.Publisher) *CreateOrderUseCase {
	return &CreateOrderUseCase{publisher: publisher}
}

func (u *CreateOrderUseCase) Execute(ctx context.Context, dto CreateOrderDto) (result *entity.OrderEntity, err error) {
	order := entity.NewOrderEntity()

	for _, item := range dto.Items {
		fakeProductName := fmt.Sprintf("Product %s", item.ProductId)
		fakeProductPrice := 10.50

		orderItem := entity.NewOrderItemEntity(fakeProductName, fakeProductPrice, item.Qtd)

		order.AddItem(orderItem)
	}

	var eventItems []event.OrderItem

	for _, item := range order.Items() {
		eventItems = append(eventItems, event.OrderItem{
			ProductName: item.ProductName(),
			Quantity:    item.Quantity(),
			TotalPrice:  item.TotalPrice(),
		})
	}

	err = u.publisher.Publish(ctx, event.OrderCreatedEvent{
		Id:         order.Id(),
		Items:      eventItems,
		TotalPrice: order.TotalPrice(),
		Status:     order.Status(),
	})

	if err != nil {
		return
	}

	return order, nil
}
