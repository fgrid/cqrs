package main

import "github.com/fgrid/uuid"

type CreateInventoryItem struct {
	InventoryItemID uuid.UUID
	Name            string
}
