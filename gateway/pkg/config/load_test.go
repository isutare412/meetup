package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFromFile(t *testing.T) {
	cfg, err := Load("../../configs/local/config.yaml")
	require.NoError(t, err)

	require.NotNil(t, cfg.Postgres)
	assert.NotEmpty(t, cfg.Postgres.Host)
}

func TestOverrideFromEnv(t *testing.T) {
	const (
		postgresPort = 15932
	)

	t.Setenv("GATEWAY_POSTGRES_PORT", strconv.Itoa(postgresPort))

	cfg, err := Load("../../configs/local/config.yaml")
	require.NoError(t, err)

	require.NotNil(t, cfg.Postgres)
	assert.Equal(t, postgresPort, cfg.Postgres.Port)
}
