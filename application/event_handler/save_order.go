package eventhandler

import (
	"order/application/repository"
	"order/domain/entity"
	"order/event"
)

type SaveOrder struct {
	orderRepository repository.OrderRepository
}

func NewSaveOrder(orderRepository repository.OrderRepository) event.EventHandler {
	return &SaveOrder{
		orderRepository: orderRepository,
	}
}

func (s *SaveOrder) Handle(evt event.Event) error {
	order := entity.ReplayOrder([]event.Event{evt})
	return s.orderRepository.Save(order)
}
