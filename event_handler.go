package cqrs

type EventHandler interface {
	Handle(e Event)
}
