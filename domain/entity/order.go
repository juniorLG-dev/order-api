package entity

import (
	appEvents "order/application/events"
	"order/domain/events"
	"order/domain/vo"
	"order/event"
	"time"
)

type Order struct {
	AggregateRoot

	ID              string
	Email           string
	Name            string
	Quantity        vo.Quantity
	Price           float64
	PaymentMethod   vo.Payment
	Location        vo.Location
	ProductID       string
	DateInformation vo.Date
}

func NewOrder(
	id, email, name, paymentMethodValue, productID string,
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
	order := Order{
		ID:            id,
		Email:         email,
		Name:          name,
		Quantity:      *quantity,
		Price:         price,
		PaymentMethod: *paymentMethod,
		Location:      location,
		ProductID:     productID,
		DateInformation: vo.Date{
			CreatedAt: time.Now().String(),
		},
	}

	order.RecordEvent(
		events.OrderPlaced{
			ID:            order.ID,
			Name:          order.Name,
			Quantity:      order.Quantity,
			Price:         order.Price,
			PaymentMethod: order.PaymentMethod,
			Location:      order.Location,
			ProductID:     order.ProductID,
		},
	)
	return &order, nil
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
		o.Email = evt.Email
		o.Name = evt.Name
		o.Quantity = evt.Quantity
		o.Price = evt.Price
		o.PaymentMethod = evt.PaymentMethod
		o.Location = evt.Location
		o.ProductID = evt.ProductID

	case appEvents.EmailSent:
		o.ID = evt.ID
		o.Email = evt.Email
		o.Name = evt.Name
		o.Quantity = evt.Quantity
		o.Price = evt.Price
		o.PaymentMethod = evt.PaymentMethod
		o.Location = evt.Location
		o.ProductID = evt.ProductID
	}
}
