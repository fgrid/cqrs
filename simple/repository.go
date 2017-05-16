package simple

import (
	"github.com/fgrid/cqrs"
	"github.com/fgrid/uuid"
)

// Repository to access aggregates
type Repository struct {
	es cqrs.EventStore
}

// NewRepository backed by the given event store (event sourced aggregates)
func NewRepository(es cqrs.EventStore) *Repository {
	return &Repository{es: es}
}

// Save an AggregateRoot with an expected Version. Will return an error if actual version of the AggregateRoot
// differs from the expected version.
func (r *Repository) Save(aggregate *cqrs.AggregateRoot, expectedVersion uint64) error {
	return nil
}

// GetByID get an Aggregate by it's ID.
func (r *Repository) GetByID(ID uuid.UUID) (*cqrs.AggregateRoot, error) {
	return nil, nil
}
