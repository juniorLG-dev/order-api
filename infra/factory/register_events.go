package factory

import (
	"errors"
	"order/event"
)

var registerEvents = make(map[string]func() event.Event)

func RegisterEvent(id string, event func() event.Event) error {
	if _, ok := registerEvents[id]; ok {
		return errors.New("this event has already been registered")
	}
	registerEvents[id] = event
	return nil
}

func GetEvent(id string) func() event.Event {
	return registerEvents[id]
}
