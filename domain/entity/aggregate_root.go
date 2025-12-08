package entity

import "order/event"

type AggregateRoot struct {
	events []event.Event
}

func (ar *AggregateRoot) RecordEvent(event event.Event) {
	ar.events = append(ar.events, event)
}

func (ar *AggregateRoot) PullEvents() []event.Event {
	events := ar.events
	ar.events = nil
	return events
}
