package domain

import "time"

type ID string
type Email string
type Password string
type EnabledAt time.Time
type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt *time.Time
type Version int
