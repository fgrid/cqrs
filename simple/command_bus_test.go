package simple_test

import (
	"testing"
	"time"

	"github.com/fgrid/cqrs"
	"github.com/fgrid/cqrs/simple"
)

type Command struct {
	// Test command
}

func TestCommandBus_deliver(t *testing.T) {
	cb := simple.NewCommandBus(1)
	answer := make(chan bool)
	defer close(answer)

	cb.RegisterFunc("*simple_test.Command", func(c cqrs.Command) {
		answer <- true
	})

	cb.Send(new(Command))

	select {
	case <-answer:
		t.Logf("command handler called")
	case <-time.After(1 * time.Millisecond):
		t.Errorf("timed out after %s", 1*time.Millisecond)
	}
}

func TestCommandBus_deliverUnknownCommand(t *testing.T) {
	cb := simple.NewCommandBus(1)
	answer := make(chan bool)
	defer close(answer)

	cb.RegisterFunc("*simple_test.OtherCommand", func(c cqrs.Command) {
		answer <- true
	})

	cb.Send(new(Command))

	select {
	case <-answer:
		t.Errorf("wrong command handler called")
	case <-time.After(1 * time.Millisecond):
		t.Logf("timed out after %s", 1*time.Millisecond)
	}
}
