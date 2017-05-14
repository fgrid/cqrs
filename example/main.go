package main

import "github.com/fgrid/cqrs/simple"

func main() {
	eventBus := simple.NewEventBus()
	eventStore := simple.NewEventStore(eventBus)
	repository := simple.NewRepository(eventStore)
	commandBus := simple.NewCommandBus()
	commands := NewInventoryCommandHandlers(repository)
	commandBus.Register("CreateInventoryItem", commands.CreateInventoryItem)

	detail := NewInventoryItemDetailView()
	eventBus.Register("InventoryItemCreated", detail.OnInventoryItemCreated)
}
