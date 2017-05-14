package cqrs

import "github.com/fgrid/uuid"

type EventStore interface {
	Store(aggregateID uuid.UUID, events []Event, expectedVersion uint64) error
	GetEventsForAggregate(aggregateID uuid.UUID) (history []Event, err error)
}
