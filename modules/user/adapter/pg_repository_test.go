package adapter_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"

	"example/modules/user/adapter"
	"example/modules/user/domain"
	"example/modules/user/port/inbound/feature"
	"example/test/container"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	container *container.PostgresContainer
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	suite.container = container.NewPostgresContainer(suite.T(), ctx)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	ctx := context.Background()
	if suite.container != nil && suite.container.Container != nil {
		if err := suite.container.Container.Terminate(ctx); err != nil {
			suite.T().Logf("Error terminating PostgreSQL container: %v", err)
		}
	}
}

func (suite *UserRepositoryTestSuite) TestCreateUser_Success() {
	// Arrange
	ctx := context.Background()
	db, err := sql.Open("postgres", suite.container.ConnStr)
	if err != nil {
		suite.T().Fatalf("Error opening SQL connection: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name TEXT,
			email TEXT,
			password TEXT,
			enabled_at TIMESTAMP,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP,
			version INT
		)
	`)
	if err != nil {
		suite.T().Fatalf("Error creating users table: %v", err)
	}

	input := feature.CreateUserRepositoryInput{
		User: domain.User{
			ID:        domain.NewID(),
			Name:      "Marcelo",
			Email:     domain.Email("marcelo@email.com"),
			Password:  domain.Password("plain-password"),
			CreatedAt: domain.NewCreatedAt(),
			UpdatedAt: domain.NewUpdatedAt(),
			Version:   domain.Version(0),
		},
	}

	// Act
	repo := adapter.NewUserRepository(db)
	err = repo.CreateUser(ctx, input)

	// Assert
	suite.NoError(err)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
