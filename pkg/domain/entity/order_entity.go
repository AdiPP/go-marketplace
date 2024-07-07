package entity

import (
	"errors"
)

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type OrderEntity struct {
	id         int64
	status     string
	items      []*OrderItemEntity
	totalPrice float64
	paidValue  float64
}

func NewOrderEntity() *OrderEntity {
	return &OrderEntity{

		status: OrderStatusPending,
	}
}

func RestoreOrderEntity(id int64, status string) *OrderEntity {
	return &OrderEntity{
		id:     id,
		status: status,
	}
}

func (o *OrderEntity) AddItem(item *OrderItemEntity) {
	o.items = append(o.items, item)
	o.totalPrice += item.TotalPrice()
}

func (o *OrderEntity) SetId(id int64) {
	o.id = id
}

func (o *OrderEntity) Id() int64 {
	return o.id
}

func (o *OrderEntity) SetStatus(status string) {
	o.status = status
}

func (o *OrderEntity) Status() string {
	return o.status
}

func (o *OrderEntity) Items() []*OrderItemEntity {
	return o.items
}

func (o *OrderEntity) SetTotalPrice(totalPrice float64) {
	o.totalPrice = totalPrice
}

func (o *OrderEntity) TotalPrice() float64 {
	return o.totalPrice
}

func (o *OrderEntity) SetPaidValue(paidValue float64) {
	o.paidValue = paidValue
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
