package events

import (
	"order/domain/vo"
	"order/event"
	"order/infra/factory"
)

type EmailSent struct {
	ID            string
	Email         string
	Name          string
	Quantity      vo.Quantity
	Price         float64
	PaymentMethod vo.Payment
	Location      vo.Location
	ProductID     string
}

func (EmailSent) GetName() string {
	return "EmailSent"
}

func init() {
	factory.RegisterEvent(EmailSent{}.GetName(), func() event.Event {
		return &EmailSent{}
	})
}
