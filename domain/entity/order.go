package entity

import (
	"order/domain/events"
	"order/domain/vo"
	"order/event"
	"time"
)

type Order struct {
	ID              string
	Name            string
	Quantity        vo.Quantity
	Price           float64
	PaymentMethod   vo.Payment
	Location        vo.Location
	ProductID       string
	DateInformation vo.Date
}

func NewOrder(
	id, name, paymentMethodValue, productID string,
	quantityValue int,
	price float64,
	location vo.Location,
) (*Order, error) {
	quantity, err := vo.NewQuantity(quantityValue)
	if err != nil {
		return nil, err
	}
	paymentMethod, err := vo.NewPayment(paymentMethodValue)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:            id,
		Name:          name,
		Quantity:      *quantity,
		Price:         price,
		PaymentMethod: *paymentMethod,
		Location:      location,
		ProductID:     productID,
		DateInformation: vo.Date{
			CreatedAt: time.Now().String(),
		},
	}, nil
}

func ReplayOrder(events []event.Event) Order {
	o := Order{}
	for _, event := range events {
		o.apply(event)
	}
	return o
}

func (o *Order) apply(e event.Event) {
	switch evt := e.(type) {
	case events.OrderPlaced:
		o.ID = evt.ID
		o.Name = evt.Name
		o.Quantity = evt.Quantity
		o.Price = evt.Price
		o.PaymentMethod = evt.PaymentMethod
		o.Location = evt.Location
		o.ProductID = evt.ProductID
	}
}
