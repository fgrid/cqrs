package example

import "github.com/fgrid/cqrs"

type CommandHandlers struct {
}

func NewInventoryCommandHandlers(r cqrs.Repository) *CommandHandlers {
	return &CommandHandlers{}
}

func (ch *CommandHandlers) Register(cb cqrs.CommandBus) {
	cb.RegisterFunc("*example.CreateInventoryItem", ch.CreateInventoryItem)
}

func (ch *CommandHandlers) CreateInventoryItem(c cqrs.Command) {
}
