package usecase

import (
	"context"
	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/AdiPP/go-marketplace/pkg/domain/event"
	"github.com/AdiPP/go-marketplace/pkg/domain/queue"
	"github.com/AdiPP/go-marketplace/pkg/usecase/repository"
)

type CreateOrderUseCase struct {
	publisher         queue.Publisher
	productRepository repository.ProductRepository
	orderRepository   repository.OrderRepository
}

func NewCreateOrderUseCase(publisher queue.Publisher, productRepository repository.ProductRepository, orderRepository repository.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{publisher: publisher, productRepository: productRepository, orderRepository: orderRepository}
}

func (u *CreateOrderUseCase) Execute(ctx context.Context, dto CreateOrderDto) (result *entity.OrderEntity, err error) {
	order := entity.NewOrderEntity()

	for _, item := range dto.Items {
		product, err := u.productRepository.FindById(item.ProductId)

		if err != nil {
			return nil, err
		}

		orderItem := entity.NewOrderItemEntity(product.ProductName, product.ProductPrice, item.Qtd)

		order.AddItem(orderItem)
	}

	order, err = u.orderRepository.Save(order)

	if err != nil {
		return
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
