package main

import "github.com/fgrid/cqrs/simple"

func main() {
	eventBus := simple.NewEventBus(2)
	eventStore := simple.NewEventStore(eventBus)
	repository := simple.NewRepository(eventStore)
	commandBus := simple.NewCommandBus(2)
	commands := NewInventoryCommandHandlers(repository)
	commandBus.RegisterFunc("*CreateInventoryItem", commands.CreateInventoryItem)

	detail := NewInventoryItemDetailView()
	eventBus.RegisterFunc("InventoryItemCreated", detail.OnInventoryItemCreated)
}
