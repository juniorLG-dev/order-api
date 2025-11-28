package event

import "order/common/event"

type EventEnvelope struct {
	ID      string
	Payload event.Event
}
