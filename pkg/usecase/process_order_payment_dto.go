package usecase

type Order struct {
	Id         string
	Items      []OrderItem
	TotalPrice float64
	Status     string
}

type OrderItem struct {
	ProductName string
	Quantity    int
	TotalPrice  float64
}

type ProcessOrderPaymentDto struct {
	Order
}
