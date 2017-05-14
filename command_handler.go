package cqrs

type CommandHandler interface {
	Handle(c Command)
}

type CommandHandlerFunc func(c Command)
