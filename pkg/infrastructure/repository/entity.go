package repository

import (
	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
)

type orderEntity struct {
	Id         int64   `db:"id" goqu:"skipinsert"`
	Status     string  `db:"status"`
	TotalPrice float64 `db:"total_price"`
	PaidValue  float64 `db:"paid_value"`
}

func newOrderEntity(entity *entity.OrderEntity) orderEntity {
	o := orderEntity{
		Id:         entity.Id(),
		Status:     entity.Status(),
		TotalPrice: entity.TotalPrice(),
		PaidValue:  entity.PaidValue(),
	}

	return o
}

func (o *orderEntity) toEntity() *entity.OrderEntity {
	e := &entity.OrderEntity{}
	e.SetId(o.Id)
	e.SetStatus(o.Status)
	e.SetTotalPrice(o.TotalPrice)
	e.SetPaidValue(o.PaidValue)

	return e
}

type orderItemEntity struct {
	Id           int64   `db:"id" goqu:"skipinsert"`
	OrderId      int64   `db:"order_id"`
	ProductId    int64   `db:"product_id"`
	ProductName  string  `db:"product_name"`
	ProductPrice float64 `db:"product_price"`
	Quantity     int64   `db:"quantity"`
}

type orderItemEntities []orderItemEntity

func newOrderItemEntities(e *entity.OrderEntity) orderItemEntities {
	o := orderItemEntities{}

	for _, v := range e.Items() {
		o = append(o, orderItemEntity{
			Id:           v.Id(),
			OrderId:      v.OrderId(),
			ProductId:    v.ProductId(),
			ProductName:  v.ProductName(),
			ProductPrice: v.ProductPrice(),
			Quantity:     v.Quantity(),
		})
	}

	return o
}

func (o orderItemEntities) setOrderId(orderId int64) orderItemEntities {
	r := orderItemEntities{}

	for _, v := range o {
		v.OrderId = orderId
		r = append(r, v)
	}

	return r
}

func (o orderItemEntities) toEntity() []*entity.OrderItemEntity {
	eArr := []*entity.OrderItemEntity{}

	for _, v := range o {
		e := &entity.OrderItemEntity{}
		e.SetId(v.Id)
		e.SetOrderId(v.OrderId)
		e.SetProductId(v.ProductId)
		e.SetProductName(v.ProductName)
		e.SetProductPrice(v.ProductPrice)
		e.SetQuantity(v.Quantity)

		eArr = append(eArr, e)
	}

	return eArr
}
