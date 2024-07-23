package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"example/modules/user/domain"
	"example/modules/user/domain/usecase"
	"example/modules/user/port/inbound/feature"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	usecase    feature.CreateUserUseCase
	repoMock   *feature.MockCreateUserRepository
	hasherMock *feature.MockPasswordHasher
}

func (suite *CreateUserUseCaseTestSuite) SetupSuite() {
	suite.repoMock = new(feature.MockCreateUserRepository)
	suite.hasherMock = new(feature.MockPasswordHasher)

	suite.usecase = usecase.NewCreateUserUseCase(suite.repoMock, suite.hasherMock)
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Success() {
	// Arrange
	input := feature.CreateUserUseCaseInput{
		Name:     "Marcelo",
		Email:    "marcelo@email.com",
		Password: "plain-password",
	}

	hashedPassword := "hashed-password"

	expectedUser := feature.CreateUserRepositoryInput{
		User: domain.User{
			Name:     input.Name,
			Email:    domain.Email(input.Email),
			Password: domain.Password(hashedPassword),
			Version:  domain.NewVersion(),
		},
	}

	// Setup mock expectations
	suite.repoMock.On("CreateUser", context.Background(), mock.MatchedBy(func(args feature.CreateUserRepositoryInput) bool {
		user := args.User
		expected := expectedUser.User

		return user.Name == expected.Name &&
			user.Email == expected.Email &&
			user.Password == expected.Password &&
			user.Version == expected.Version
	})).Return(nil)
	suite.hasherMock.On("Hash", input.Password).Return(hashedPassword, nil)

	// Act
	output, err := suite.usecase.Execute(context.Background(), input)

	// Assert
	suite.NoError(err)
	suite.Equal(input.Name, output.User.Name, "Name should be equal")
	suite.Equal(input.Email, string(output.User.Email), "Email should be equal")
	suite.Equal(hashedPassword, string(output.User.Password), "Password should be equal")
	suite.NotNil(output.User.CreatedAt, "CreatedAt should not be nil")
	suite.NotNil(output.User.UpdatedAt, "UpdatedAt should not be nil")
	suite.Equal(output.User.Version, domain.Version(0), "Version should be greater than 0")

	// Verify mock expectations
	suite.hasherMock.AssertExpectations(suite.T())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_FailHashPassword() {
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_FailCreateUser() {
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
