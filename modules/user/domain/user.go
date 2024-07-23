package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	EnabledAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
