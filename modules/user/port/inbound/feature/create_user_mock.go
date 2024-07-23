package feature

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCreateUserRepository struct {
	mock.Mock
}

func (m *MockCreateUserRepository) CreateUser(ctx context.Context, input CreateUserRepositoryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHasher) Compare(data, encodedHash string) (bool, error) {
	args := m.Called(data, encodedHash)
	return args.Bool(0), args.Error(1)
}
