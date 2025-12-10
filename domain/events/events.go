package events

import (
	"order/domain/vo"
	"order/event"
	"order/infra/factory"
)

type OrderPlaced struct {
	ID            string
	Email         string
	Name          string
	Quantity      vo.Quantity
	Price         float64
	PaymentMethod vo.Payment
	Location      vo.Location
	ProductID     string
	CreatedAt     string
}

func (OrderPlaced) GetName() string {
	return "OrderPlaced"
}

func init() {
	factory.RegisterEvent(OrderPlaced{}.GetName(), func() event.Event {
		return &OrderPlaced{}
	})
}
