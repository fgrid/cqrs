package main

import "github.com/fgrid/uuid"

type InventoryItemCreated struct {
	ID   uuid.UUID
	Name string
}
