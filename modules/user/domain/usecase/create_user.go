package usecase

import (
	"context"
	"fmt"

	"example/modules/user/domain"
	"example/modules/user/port/inbound/feature"
)

type CreateUserUseCase struct {
	repository feature.CreateUserRepository
	hasher     feature.PasswordHasher
}

func NewCreateUserUseCase(repository feature.CreateUserRepository, hasher feature.PasswordHasher) feature.CreateUserUseCase {
	return &CreateUserUseCase{repository: repository, hasher: hasher}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input feature.CreateUserUseCaseInput) (feature.CreateUserUseCaseOutput, error) {
	hashedPassword, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return feature.CreateUserUseCaseOutput{}, domain.GetErrUserPasswordHashFailed(err)
	}

	user := domain.User{
		ID:        domain.NewID(),
		Name:      input.Name,
		Email:     domain.Email(input.Email),
		Password:  domain.Password(hashedPassword),
		CreatedAt: domain.NewCreatedAt(),
		UpdatedAt: domain.NewUpdatedAt(),
		Version:   domain.NewVersion(),
	}

	fmt.Println("user", user)

	err = uc.repository.CreateUser(ctx, feature.CreateUserRepositoryInput{User: user})
	if err != nil {
		return feature.CreateUserUseCaseOutput{}, domain.GetErrUserPersistNewUserFailed(err)
	}

	return feature.CreateUserUseCaseOutput{User: user}, nil
}
