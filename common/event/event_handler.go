package event

type EventHandler interface {
	Handle(Event) error
}
