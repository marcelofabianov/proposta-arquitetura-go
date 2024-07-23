package usecase_test

import (
	"context"
	"testing"

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

	// Setup mock expectations
	suite.hasherMock.On("Hash", input.Password).Return(hashedPassword, nil)
	suite.repoMock.On("CreateUser", context.Background(), feature.CreateUserRepositoryInput{
		User: domain.User{
			ID:        domain.NewID(),
			Name:      input.Name,
			Email:     domain.Email(input.Email),
			Password:  domain.Password(hashedPassword),
			EnabledAt: domain.EnabledAt(nil),
			CreatedAt: domain.NewCreatedAt(),
			UpdatedAt: domain.NewUpdatedAt(),
			Version:   domain.NewVersion(),
		},
	}).Return(nil)

	// Act
	output, err := suite.usecase.Execute(context.Background(), input)

	// Assert
	suite.NoError(err)
	suite.Equal(input.Name, output.User.Name, "Name should be equal")
	suite.Equal(input.Email, string(output.User.Email), "Email should be equal")
	suite.Equal(hashedPassword, string(output.User.Password), "Password should be equal")
	suite.NotNil(output.User.EnabledAt, "EnabledAt should not be nil")
	suite.NotNil(output.User.CreatedAt, "CreatedAt should not be nil")
	suite.NotNil(output.User.UpdatedAt, "UpdatedAt should not be nil")
	suite.Greater(output.User.Version, int64(0), "Version should be greater than 0")

	// Verify mock expectations
	suite.hasherMock.AssertExpectations(suite.T())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Error_UserPasswordHashFailed() {
	//...

	suite.Assert().Fail("TestCreateUser_Error_UserPasswordHashFailed not implemented yet")
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Error_UserPersistNewUserFailed() {
	//...

	suite.Assert().Fail("TestCreateUser_Error_UserPersistNewUserFailed not implemented yet")
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
