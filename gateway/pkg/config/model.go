package config

import "time"

type Config struct {
	Timeout  *TimeoutConfig  `yaml:"timeout"`
	Logger   *LoggerConfig   `yaml:"logger"`
	Server   *ServerConfig   `yaml:"server"`
	Postgres *PostgresConfig `yaml:"postgres"`
}

type TimeoutConfig struct {
	Startup  time.Duration `yaml:"startup"`
	Shutdown time.Duration `yaml:"shutdown"`
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
}

type ServerConfig struct {
	HTTP *HTTPServerConfig `yaml:"http"`
}

type HTTPServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
