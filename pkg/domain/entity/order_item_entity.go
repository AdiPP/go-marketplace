package entity

type OrderItemEntity struct {
	productName  string
	productPrice float64
	quantity     int
}

func NewOrderItemEntity(productName string, productPrice float64, quantity int) *OrderItemEntity {
	return &OrderItemEntity{productName: productName, productPrice: productPrice, quantity: quantity}
}

func (o *OrderItemEntity) ProductName() string {
	return o.productName
}

func (o *OrderItemEntity) ProductPrice() float64 {
	return o.productPrice
}

func (o *OrderItemEntity) Quantity() int {
	return o.quantity
}

func (o *OrderItemEntity) TotalPrice() float64 {
	return o.productPrice * float64(o.quantity)
}
