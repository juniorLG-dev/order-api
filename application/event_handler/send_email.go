package eventhandler

import (
	"fmt"
	"order/application/events"
	"order/application/smtp"
	"order/domain/entity"
	"order/event"
)

type SendEmail struct {
	orderSMTP smtp.OrderSMTP
	eventBus  event.EventBus
}

func NewSendEmail(
	orderSMTP smtp.OrderSMTP,
	eventBus event.EventBus,
) event.EventHandler {
	return &SendEmail{
		orderSMTP: orderSMTP,
		eventBus:  eventBus,
	}
}

func (s *SendEmail) Handle(evt event.Event) error {
	order := entity.ReplayOrder([]event.Event{evt})
	err := s.orderSMTP.SendEmail(
		order.Email,
		"Order processed successfully!",
		fmt.Sprintf(`
			<h1>The purchase order for the product: test was approved</h1>
			<br>
			<p>It was shipped to the following location: Country - <strong>%s</strong>, State - <strong>%s</strong>, Postal Code(CEP) - <strong>%s</strong></p>
		`, order.Location.Country, order.Location.State, order.Location.CEP.Value),
	)
	if err != nil {
		return err
	}
	return s.eventBus.Publish(events.EmailSent{
		ID:            order.ID,
		Email:         order.Email,
		Name:          order.Name,
		Quantity:      order.Quantity,
		Price:         order.Price,
		PaymentMethod: order.PaymentMethod,
		Location:      order.Location,
		ProductID:     order.ProductID,
	})
}
