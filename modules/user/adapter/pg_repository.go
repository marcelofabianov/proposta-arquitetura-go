package adapter

import (
	"context"
	"database/sql"

	"example/modules/user/port/inbound"
	"example/modules/user/port/inbound/feature"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) inbound.UserRepositoryInbound {
	return &UserRepository{Db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, input feature.CreateUserRepositoryInput) error {
	query := `INSERT INTO users (id, name, email, password, enabled_at, created_at, updated_at, version) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	result, err := r.Db.ExecContext(
		ctx,
		query,
		input.User.ID,
		input.User.Name,
		input.User.Email,
		input.User.Password,
		input.User.EnabledAt,
		input.User.CreatedAt,
		input.User.UpdatedAt,
		0, // version
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	return err
}
