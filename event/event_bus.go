package event

type EventBus interface {
	Publish(events ...Event) error
	Subscribe(eventName string, evt EventHandler) error
}
