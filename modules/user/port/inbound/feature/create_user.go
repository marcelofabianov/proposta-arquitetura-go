package feature

import (
	"context"

	"example/modules/user/domain"
)

// PKG

type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

// Repository

type CreateUserRepositoryInput struct {
	User domain.User
}

type CreateUserRepository interface {
	CreateUser(ctx context.Context, input CreateUserRepositoryInput) error
}

type UserRepository interface {
	CreateUserRepository
}

// UseCase

type CreateUserUseCaseInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserUseCaseOutput struct {
	User domain.User
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserUseCaseInput) (CreateUserUseCaseOutput, error)
}

// Service

type CreateUserServiceInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserServiceOutput struct {
	ID        string
	Name      string
	Email     string
	EnabledAt string
	CreatedAt string
	UpdatedAt string
}

// Service / Inbound
type CreateUserServiceInboundInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserServiceInboundOutput struct {
	ID        string
	Name      string
	Email     string
	EnabledAt string
	CreatedAt string
	UpdatedAt string
}

type CreateUserServiceInbound interface {
	CreateUser(ctx context.Context, req CreateUserServiceInboundInput) (*CreateUserServiceInboundOutput, error)
}

// Request

type CreateUserRequest struct {
	Name     string `json:"name" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

// Presenter

type CreateUserPresenter struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	EnabledAt string `json:"enabled_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
