package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

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
	//...

	suite.Fail("TestCreateUser_Success not implemented yet")
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Error_UserPasswordHashFailed() {
	//...

	suite.Fail("TestCreateUser_Error_UserPasswordHashFailed not implemented yet")
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUser_Error_UserPersistNewUserFailed() {
	//...

	suite.Fail("TestCreateUser_Error_UserPersistNewUserFailed not implemented yet")
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
