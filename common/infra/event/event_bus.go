package event

import (
	"encoding/json"
	"order/common/event"

	"github.com/streadway/amqp"
)

type EventBus struct {
	channel *amqp.Channel
	name    string
}

func NewEventBus(channel *amqp.Channel, name string) *EventBus {
	return &EventBus{
		channel: channel,
		name:    name,
	}
}

func (e *EventBus) Publish(events ...event.Event) error {
	err := e.channel.ExchangeDeclare(
		e.name,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for _, event := range events {
		eventPayload, err := json.Marshal(event)
		if err != nil {
			return nil
		}
		err = e.channel.Publish(
			e.name,
			event.GetName(),
			false,
			false,
			amqp.Publishing{
				ContentType: "enconding/json",
				Body:        eventPayload,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EventBus) Subscribe(eventName string, evt event.EventHandler) error {
	q, err := e.channel.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	err = e.channel.QueueBind(
		q.Name,
		eventName,
		e.name,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := e.channel.Consume(
		q.Name,
		"order",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			var e event.Event
			if err := json.Unmarshal(msg.Body, &e); err == nil {
				evt.Handle(e)
			}
		}
	}()
	return nil
}
