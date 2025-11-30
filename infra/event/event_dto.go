package event

import "order/event"

type EventEnvelope struct {
	ID      string
	Payload event.Event
}
