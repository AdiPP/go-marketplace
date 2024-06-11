package event

type OrderCreatedEvent struct {
	Id         string
	Items      []OrderItem
	TotalPrice float64
	Status     string
}

func (o OrderCreatedEvent) GetType() string {
	return "OrderCreatedEvent"
}

type OrderItem struct {
	ProductName string
	Quantity    int
	TotalPrice  float64
}
