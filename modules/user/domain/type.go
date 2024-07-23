package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID string
type Email string
type Password string
type EnabledAt *time.Time
type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt *time.Time
type Version int64

func NewID() ID {
	return ID(uuid.New().String())
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func NewVersion() Version {
	return Version(0)
}
