package command

import (
	"order/domain/entity"
	"order/domain/vo"
	"order/event"
)

type PlaceOrder struct {
	eventBus event.EventBus
}

func NewPlaceOrder(
	eventBus event.EventBus,
) *PlaceOrder {
	return &PlaceOrder{
		eventBus: eventBus,
	}
}

func (p *PlaceOrder) Run(input PlaceOrderInput) error {
	order, err := entity.NewOrder(
		input.ID,
		input.Name,
		input.PaymentMethod,
		input.ProductID,
		input.Quantity,
		input.Price,
		input.Location,
	)
	if err != nil {
		return err
	}
	p.eventBus.Publish(order.PullEvents()...)
	return nil
}

type PlaceOrderInput struct {
	ID            string
	Name          string      `json:"name"`
	Quantity      int         `json:"quantity"`
	Price         float64     `json:"price"`
	PaymentMethod string      `json:"payment_method"`
	Location      vo.Location `json:"location"`
	ProductID     string      `json:"product_id"`
}
