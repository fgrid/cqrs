package cqrs

// EventBus will route a published event to every handler who is registered to this kind of event.
type EventBus interface {

	// Register an EventHandler for the given kind of event
	Register(kind string, eh EventHandler)

	// Publish an event
	Publish(e Event)
}
