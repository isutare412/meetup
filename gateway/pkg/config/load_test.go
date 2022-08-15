package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFromFile(t *testing.T) {
	cfg, err := Load("sample/config.yaml")
	require.NoError(t, err)

	require.NotNil(t, cfg.Postgres)
	assert.NotEmpty(t, cfg.Postgres.Addr)
}

func TestOverrideFromEnv(t *testing.T) {
	const (
		postgresAddr = "localhost:15932"
	)

	t.Setenv("GATEWAY_POSTGRES_ADDR", postgresAddr)

	cfg, err := Load("sample/config.yaml")
	require.NoError(t, err)

	require.NotNil(t, cfg.Postgres)
	assert.Equal(t, cfg.Postgres.Addr, postgresAddr)
}
