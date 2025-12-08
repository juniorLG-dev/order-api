package command_test

import (
	"order/application/command"
	"order/domain/events"
	"order/domain/vo"
	"order/infra/event"
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	eventBus := event.NewEventBusMocked()
	expectedEvent := events.OrderPlaced{}
	c := command.NewPlaceOrder(eventBus)
	c.Run(command.PlaceOrderInput{
		ID:            "id-123",
		Name:          "test",
		Quantity:      1,
		Price:         100,
		PaymentMethod: "PIX",
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
