package cqrs

type CommandHandler interface {
	Handle(c Command)
}
