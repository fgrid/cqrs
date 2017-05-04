package cqrs

// CommandSender will send a command to the bounded context
type CommandSender interface {
	Send(c Command) error
}
