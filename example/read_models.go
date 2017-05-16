package main

import "github.com/fgrid/cqrs"

type InventoryItemDetailView struct {
}

func NewInventoryItemDetailView() *InventoryItemDetailView {
	return &InventoryItemDetailView{}
}

func (v *InventoryItemDetailView) OnInventoryItemCreated(e cqrs.Event) {

}
