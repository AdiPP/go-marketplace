package event

import "time"

type OrderPaidEvent struct {
	OrderId     int64
	PaidValue   float64
	PaymentDate time.Time
}

func (o OrderPaidEvent) GetType() string {
	return "OrderPaidEvent"
}
