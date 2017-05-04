package cqrs

type EventPublisher interface {
	Publish(e Event)
}
