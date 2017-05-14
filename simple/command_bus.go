package simple

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fgrid/cqrs"
)

// CommandBus will route a sent command to one of the registered handlers.
// This simple command bus will route to command handlers dwelling within the same process only.
type CommandBus struct {
	bus     chan cqrs.Command
	handler map[string][]cqrs.CommandHandlerFunc
}

// NewCommandBus will return a new simple command bus.
func NewCommandBus(workers int) *CommandBus {
	cb := &CommandBus{
		bus:     make(chan cqrs.Command, 100),
		handler: make(map[string][]cqrs.CommandHandlerFunc),
	}
	for i := 0; i < workers; i++ {
		go cb.deliver(i)
	}
	return cb
}

// Register a CommandHandler to the given kind of command
func (cb *CommandBus) Register(k string, h cqrs.CommandHandler) {
	cb.RegisterFunc(k, h.Handle)
}

// RegisterFunc will register the given function as command handler for the given kind of command
func (cb *CommandBus) RegisterFunc(k string, chf cqrs.CommandHandlerFunc) {
	cb.handler[k] = append(cb.handler[k], chf)
}

// Send a command
func (cb *CommandBus) Send(c cqrs.Command) {
	cb.bus <- c
}

func (cb *CommandBus) deliver(worker int) {
	for c := range cb.bus {
		kind := fmt.Sprintf("%T", c)
		if handlers, found := cb.handler[kind]; !found {
			log.Printf("[worker-%03d] no handler for commands of kind %q", worker, kind)
		} else {
			start := time.Now()
			handlers[rand.Intn(len(handlers))](c)
			log.Printf("[worker-%03d] handling command %q took %s", worker, kind, time.Since(start))
		}
	}
}
