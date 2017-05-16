package simple

import (
	"fmt"
	"log"

	"github.com/fgrid/cqrs"
)

type EventBus struct {
	bus     chan cqrs.Command
	handler map[string][]cqrs.EventHandlerFunc
}

func NewEventBus(workers int) *EventBus {
	eb := &EventBus{
		bus:     make(chan cqrs.Command, 100),
		handler: make(map[string][]cqrs.EventHandlerFunc),
	}
	for i := 0; i < workers; i++ {
		go eb.deliver(i)
	}
	return eb
}

func (eb *EventBus) Register(k string, h cqrs.EventHandler) {
	eb.RegisterFunc(k, h.Handle)
}

func (eb *EventBus) RegisterFunc(k string, ehf cqrs.EventHandlerFunc) {
	eb.handler[k] = append(eb.handler[k], ehf)
}

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
