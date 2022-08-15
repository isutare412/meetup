package config

type Config struct {
	Postgres *PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Addr string `yaml:"addr"`
}
