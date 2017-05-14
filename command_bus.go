package cqrs

// CommandBus will route a sent command to one of the registered handlers.
type CommandBus interface {

	// Register a CommandHandler to the given kind of command
	Register(kind string, ch CommandHandler)

	// RegisterFunc will register the given function as command handler for the given kind of command
	RegisterFunc(kind string, chf CommandHandlerFunc)

	// Send a command
	Send(c Command)
}
