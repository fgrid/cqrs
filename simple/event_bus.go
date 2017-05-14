package simple

import "github.com/fgrid/cqrs"

type EventBus struct {
}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (eb *EventBus) Register(k string, h cqrs.EventHandler) {
}

func (eb *EventBus) Publish(e cqrs.Event) {
}
