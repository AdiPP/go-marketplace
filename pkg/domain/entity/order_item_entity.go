package entity

type OrderItemEntity struct {
	id           int64
	orderId      int64
	productId    int64
	productName  string
	productPrice float64
	quantity     int64
}

func NewOrderItemEntity(productName string, productPrice float64, quantity int64) *OrderItemEntity {
	return &OrderItemEntity{productName: productName, productPrice: productPrice, quantity: quantity}
}

func (o *OrderItemEntity) SetId(id int64) {
	o.id = id
}

func (o *OrderItemEntity) Id() int64 {
	return o.id
}

func (o *OrderItemEntity) SetOrderId(orderId int64) {
	o.orderId = orderId
}

func (o *OrderItemEntity) OrderId() int64 {
	return o.orderId
}

func (o *OrderItemEntity) SetProductId(productId int64) {
	o.productId = productId
}

func (o *OrderItemEntity) ProductId() int64 {
	return o.productId
}

func (o *OrderItemEntity) SetProductName(productName string) {
	o.productName = productName
}

func (o *OrderItemEntity) ProductName() string {
	return o.productName
}

func (o *OrderItemEntity) SetProductPrice(productPrice float64) {
	o.productPrice = productPrice
}

func (o *OrderItemEntity) ProductPrice() float64 {
	return o.productPrice
}

func (o *OrderItemEntity) SetQuantity(quantity int64) {
	o.quantity = quantity
}

func (o *OrderItemEntity) Quantity() int64 {
	return o.quantity
}

func (o *OrderItemEntity) TotalPrice() float64 {
	return o.productPrice * float64(o.quantity)
}
