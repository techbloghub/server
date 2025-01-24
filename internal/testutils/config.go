package testutils

import (
	"testing"

	"github.com/techbloghub/server/config"
)

func NewTestConfig(t *testing.T) (*config.Config, error) {
	t.Setenv("ENV", "test")
	t.Setenv("PORT", "8081")
	t.Setenv("POSTGRES_HOST", "localhost")
	t.Setenv("POSTGRES_USER", "example-user")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "tbh-db")
	t.Setenv("POSTGRES_PORT", "5433")

	return config.NewConfig()
}
