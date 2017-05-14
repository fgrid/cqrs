package simple

import (
	"github.com/fgrid/cqrs"
	"github.com/fgrid/uuid"
)

type EventStore struct {
	p cqrs.EventPublisher
}

func NewEventStore(p cqrs.EventPublisher) *EventStore {
	return &EventStore{p: p}
}

func (es *EventStore) Store(ID uuid.UUID, events []cqrs.Event, expectedVersion uint64) error {
	return nil
}

func (es *EventStore) GetEventsForAggregate(ID uuid.UUID) (history []cqrs.Event, err error) {
	return history, nil
}
