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

func (suite *CreateUserUseCaseTestSuite) SetupTest() {
	suite.repoMock = new(feature.MockCreateUserRepository)
	suite.hasherMock = new(feature.MockPasswordHasher)
	suite.usecase = usecase.NewCreateUserUseCase(suite.repoMock, suite.hasherMock)
}

func (suite *CreateUserUseCaseTestSuite) TearDownTest() {
	suite.repoMock.AssertExpectations(suite.T())
	suite.hasherMock.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Success() {
	// Arrange
	input := feature.CreateUserUseCaseInput{
		Name:     "Marcelo",
		Email:    "marcelo@email.com",
		Password: "plain-password",
	}

	hashedPassword := "hashed-password"

	expectedUser := domain.User{
		Name:     input.Name,
		Email:    domain.Email(input.Email),
		Password: domain.Password(hashedPassword),
		Version:  domain.NewVersion(),
	}

	// Setup mock expectations
	suite.repoMock.On("CreateUser", context.Background(), mock.MatchedBy(func(args feature.CreateUserRepositoryInput) bool {
		user := args.User
		return user.Name == expectedUser.Name &&
			user.Email == expectedUser.Email &&
			user.Password == expectedUser.Password &&
			user.Version == expectedUser.Version
	})).Return(nil)
	suite.hasherMock.On("Hash", input.Password).Return(hashedPassword, nil)

	// Act
	output, err := suite.usecase.Execute(context.Background(), input)

	// Assert
	suite.NoError(err, "Expected no error")
	suite.Equal(input.Name, output.User.Name, "Name should be equal")
	suite.Equal(input.Email, string(output.User.Email), "Email should be equal")
	suite.Equal(hashedPassword, string(output.User.Password), "Password should be equal")
	suite.NotNil(output.User.CreatedAt, "CreatedAt should not be nil")
	suite.NotNil(output.User.UpdatedAt, "UpdatedAt should not be nil")
	suite.Equal(output.User.Version, domain.Version(0), "Version should be 0")
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Fail_ErrUserPasswordHashFailed() {
	// Arrange
	input := feature.CreateUserUseCaseInput{
		Name:     "Marcelo",
		Email:    "marcelo@email.com",
		Password: "plain-password",
	}

	// Setup mock expectations
	expectedError := domain.GetErrUserPasswordHashFailed(nil)
	suite.hasherMock.On("Hash", input.Password).Return("", expectedError)

	// Act
	output, err := suite.usecase.Execute(context.Background(), input)

	// Assert
	suite.Error(err)
	suite.True(domain.IsErrUserPasswordHashFailed(err), "Error should be ErrUserPasswordHashFailed")
	suite.Equal(feature.CreateUserUseCaseOutput{}, output, "Output should be empty in case of error")
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Fail_ErrUserPersistNewUserFailed() {
	// Arrange
	input := feature.CreateUserUseCaseInput{
		Name:     "Marcelo",
		Email:    "marcelo@email.com",
		Password: "plain-password",
	}

	hashedPassword := "hashed-password"
	expectedError := domain.GetErrUserPersistNewUserFailed(nil)

	// Setup mock expectations
	suite.hasherMock.On("Hash", input.Password).Return(hashedPassword, nil)
	suite.repoMock.On("CreateUser", context.Background(), mock.MatchedBy(func(args feature.CreateUserRepositoryInput) bool {
		return args.User.Name == input.Name &&
			args.User.Email == domain.Email(input.Email) &&
			args.User.Password == domain.Password(hashedPassword)
	})).Return(expectedError)

	// Act
	output, err := suite.usecase.Execute(context.Background(), input)

	// Assert
	suite.Error(err)
	suite.True(domain.IsErrUserPersistNewUserFailed(err), "Error should be ErrUserCreationFailed")
	suite.Equal(feature.CreateUserUseCaseOutput{}, output, "Output should be empty in case of error")
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
