package features

import (
	"context"

	"example/modules/user/domain"
)

// Repository

type CreateUserRepositoryInput struct {
	User domain.User
}

type CreateUserRepository interface {
	CreateUser(ctx context.Context, input CreateUserRepositoryInput) error
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
	CreateUser(ctx context.Context, input CreateUserUseCaseInput) (CreateUserUseCaseOutput, error)
}

// Service

type CreateUserServiceRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=150"`
}

type CreateUserServicePresenter struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	EnabledAt string `json:"enabled_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUserService interface {
	CreateUser(ctx context.Context, request CreateUserServiceRequest) (CreateUserServicePresenter, error)
}
