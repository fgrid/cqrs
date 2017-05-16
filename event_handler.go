package cqrs

type EventHandler interface {
	Handle(e Event)
}
type EventHandlerFunc func(e Event)
