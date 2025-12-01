package event

import (
	"order/event"
	"sync"
)

type EventBusMocked struct {
	mu     sync.Mutex
	events []event.Event
}

func NewEventBusMocked() event.EventBus {
	return &EventBusMocked{
		events: make([]event.Event, 0),
	}
}

func (e *EventBusMocked) Publish(events ...event.Event) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.events = append(e.events, events...)
	return nil
}

func (e *EventBusMocked) EventExists(evt event.Event) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, event := range e.events {
		if event.GetName() == evt.GetName() {
			return true
		}
	}
	return false
}

func (e *EventBusMocked) Subscribe(eventName string, evt event.EventHandler) error {
	return nil
}
