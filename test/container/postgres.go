package container

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type PostgresContainer struct {
	Container *postgres.PostgresContainer
	ConnStr   string
}

func NewPostgresContainer(t *testing.T, ctx context.Context) *PostgresContainer {
	container, err := postgres.Run(
		ctx,
		"postgres:16",
		postgres.WithDatabase("dbname"),
		postgres.WithUsername("username"),
		postgres.WithPassword("password"),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatal(err)
		}
	})

	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "timezone=UTC", "application_name=test")
	if err != nil {
		t.Fatal(err)
	}

	return &PostgresContainer{
		Container: container,
		ConnStr:   connStr,
	}
}
