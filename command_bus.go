package cqrs

// CommandBus will route a sent command to one of the registered handlers.
type CommandBus interface {

	// Register a CommandHandler to the given kind of command
	Register(kind string, ch CommandHandler)

	// Send a command
	Send(c Command)
}
