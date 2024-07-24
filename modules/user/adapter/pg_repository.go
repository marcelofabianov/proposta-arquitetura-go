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
	query := `INSERT INTO users (id, name, email, password, enabled_at, created_at, updated_at, version) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	result, err := r.Db.ExecContext(
		ctx,
		query,
		input.User.ID.String(),
		input.User.Name,
		input.User.Email.String(),
		input.User.Password.String(),
		input.User.EnabledAt,
		input.User.CreatedAt.String(),
		input.User.UpdatedAt.String(),
		input.User.Version.String(),
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	return err
}
