package cqrs

import "github.com/fgrid/uuid"

// Repository of AggregateRoots
type Repository interface {
	// Save an AggregateRoot with an expected Version. Will return an error if actual version of the AggregateRoot
	// differs from the expected version.
	Save(aggregate *AggregateRoot, expectedVersion uint64) error

	// GetByID get an Aggregate by it's ID.
	GetByID(ID uuid.UUID) (*AggregateRoot, error)
}
