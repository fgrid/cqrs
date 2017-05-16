package simple

import (
	"fmt"
	"log"

	"github.com/fgrid/cqrs"
)

// EventBus for transportation of domaine events
type EventBus struct {
	bus     chan cqrs.Event
	handler map[string][]cqrs.EventHandlerFunc
}

// NewEventBus creates a new event bus with given count of workers used to deliver events to subscribers.
func NewEventBus(workers int) *EventBus {
	eb := &EventBus{
		bus:     make(chan cqrs.Event, 100),
		handler: make(map[string][]cqrs.EventHandlerFunc),
	}
	for i := 0; i < workers; i++ {
		go eb.deliver(i)
	}
	return eb
}

// Register an event handler for the given kind of events.
func (eb *EventBus) Register(k string, h cqrs.EventHandler) {
	eb.RegisterFunc(k, h.Handle)
}

// RegisterFunc to handle the given kind of events.
func (eb *EventBus) RegisterFunc(k string, ehf cqrs.EventHandlerFunc) {
	eb.handler[k] = append(eb.handler[k], ehf)
}

// Publish an event.
func (eb *EventBus) Publish(e cqrs.Event) {
	eb.bus <- e
}

func (eb *EventBus) deliver(worker int) {
	for e := range eb.bus {
		kind := fmt.Sprintf("%T", e)
		if handlers, found := eb.handler[kind]; !found {
			log.Printf("[worker-%03d] no handler for events of kind %q", worker, kind)
		} else {
			for _, eh := range handlers {
				eh(e)
			}
		}
	}
}
