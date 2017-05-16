package simple_test

import (
	"testing"
	"time"

	"github.com/fgrid/cqrs"
	"github.com/fgrid/cqrs/simple"
)

type Event struct {
	// Test command
}

func TestEventBus_deliver(t *testing.T) {
	eb := simple.NewEventBus(1)
	answer := make(chan bool)
	defer close(answer)

	eb.RegisterFunc("*simple_test.Event", func(e cqrs.Event) {
		t.Logf("event handler1 called")
		answer <- true
	})
	eb.RegisterFunc("*simple_test.Event", func(e cqrs.Event) {
		t.Logf("event handler2 called")
		answer <- true
	})

	eb.Publish(new(Event))

	for count := 0; count < 2; {
		select {
		case <-answer:
			count++
		case <-time.After(1 * time.Millisecond):
			t.Errorf("timed out after %s", 1*time.Millisecond)
			return
		}
	}
}

func TestEventBus_deliverUnknownCommand(t *testing.T) {
	eb := simple.NewEventBus(1)
	answer := make(chan bool)
	defer close(answer)

	eb.RegisterFunc("*simple_test.OtherCommand", func(e cqrs.Event) {
		answer <- true
	})

	eb.Publish(new(Event))

	select {
	case <-answer:
		t.Errorf("wrong command handler called")
	case <-time.After(1 * time.Millisecond):
		t.Logf("timed out after %s", 1*time.Millisecond)
	}
}
