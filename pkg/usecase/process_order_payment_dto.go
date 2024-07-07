package usecase

type Order struct {
	Id         int64
	Items      []OrderItem
	TotalPrice float64
	Status     string
}

type OrderItem struct {
	ProductName string
	Quantity    int64
	TotalPrice  float64
}

type ProcessOrderPaymentDto struct {
	Order
}
