package event

type OrderCreatedEvent struct {
	Id         int64
	Items      []OrderItem
	TotalPrice float64
	Status     string
}

func (o OrderCreatedEvent) GetType() string {
	return "OrderCreatedEvent"
}

type OrderItem struct {
	ProductName string
	Quantity    int64
	TotalPrice  float64
}
