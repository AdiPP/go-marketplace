package entity

import (
	"errors"
	"github.com/google/uuid"
)

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type OrderEntity struct {
	id         string
	status     string
	items      []*OrderItemEntity
	totalPrice float64
	paidValue  float64
}

func NewOrderEntity() *OrderEntity {
	return &OrderEntity{
		id:     uuid.New().String(),
		status: OrderStatusPending,
	}
}

func RestoreOrderEntity(id, status string) *OrderEntity {
	return &OrderEntity{
		id:     id,
		status: status,
	}
}

func (o *OrderEntity) AddItem(item *OrderItemEntity) {
	o.items = append(o.items, item)
	o.totalPrice += item.TotalPrice()
}

func (o *OrderEntity) Id() string {
	return o.id
}

func (o *OrderEntity) Status() string {
	return o.status
}

func (o *OrderEntity) Items() []*OrderItemEntity {
	return o.items
}

func (o *OrderEntity) TotalPrice() float64 {
	return o.totalPrice
}

func (o *OrderEntity) PaidValue() float64 {
	return o.paidValue
}

func (o *OrderEntity) Pay(value float64) error {
	if value < o.totalPrice {
		return errors.New("value is lower than total price")
	}

	o.paidValue = value
	o.status = OrderStatusPaid

	return nil
}
