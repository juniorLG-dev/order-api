package eventhandler_test

import (
	eventhandler "order/application/event_handler"
	"order/application/events"
	"order/domain/vo"
	"order/infra/event"
	"order/infra/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	eventBus := event.NewEventBusMocked()
	smtp := smtp.NewSMTPMocked()
	expectedEvent := events.EmailSent{}
	e := eventhandler.NewSendEmail(smtp, eventBus)
	e.Handle(events.EmailSent{
		ID:            "id-123",
		Email:         "email@example.com",
		Name:          "test",
		Quantity:      vo.Quantity{Value: 1},
		Price:         100,
		PaymentMethod: vo.Payment{Value: "PIX"},
		Location: vo.Location{
			Country:    "Brazil",
			State:      "Rio Grande do Norte",
			City:       "Natal",
			Complement: "Rua...",
			CEP: vo.CEP{
				Value: "12345-678",
			},
		},
		ProductID: "id-456",
	})
	if !eventBus.EventExists(expectedEvent) {
		t.Fail()
	}
}
