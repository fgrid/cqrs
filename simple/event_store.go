package simple

import (
	"github.com/fgrid/cqrs"
	"github.com/fgrid/uuid"
)

// EventStore is just a simple event store implementation
type EventStore struct {
	p cqrs.EventPublisher
}

// NewEventStore creates a new simple event store that publishes the events after storing with the given publisher.
func NewEventStore(p cqrs.EventPublisher) *EventStore {
	return &EventStore{p: p}
}

// Store events for the aggregate specified by the given ID expecting the aggregate in the given version.
func (es *EventStore) Store(ID uuid.UUID, events []cqrs.Event, expectedVersion uint64) error {
	return nil
}

// GetEventsForAggregate gets the history of events for the aggregate specified by the given ID.
func (es *EventStore) GetEventsForAggregate(ID uuid.UUID) (history []cqrs.Event, err error) {
	return history, nil
}
