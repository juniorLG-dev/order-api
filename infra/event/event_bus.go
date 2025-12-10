package event

import (
	"encoding/json"
	"order/event"
	"order/infra/factory"

	"github.com/streadway/amqp"
)

type EventBus struct {
	channel *amqp.Channel
	name    string
}

func NewEventBus(channel *amqp.Channel, name string) (*EventBus, error) {
	err := channel.ExchangeDeclare(
		name,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &EventBus{
		channel: channel,
		name:    name,
	}, nil
}

func (e *EventBus) Publish(events ...event.Event) error {
	for _, event := range events {
		eventJSON, _ := json.Marshal(event)
		wrappedEvent, err := json.Marshal(
			EventEnvelope{
				ID:      event.GetName(),
				Payload: eventJSON,
			},
		)
		if err != nil {
			return err
		}
		err = e.channel.Publish(
			e.name,
			event.GetName(),
			false,
			false,
			amqp.Publishing{
				ContentType: "enconding/json",
				Body:        wrappedEvent,
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
		q.Name,
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
			var envelope EventEnvelope
			json.Unmarshal(msg.Body, &envelope)
			targetEvent := factory.GetEvent(eventName)
			event := targetEvent()
			if err := json.Unmarshal(envelope.Payload, event); err == nil {
				evt.Handle(event)
			}
		}
	}()
	return nil
}
