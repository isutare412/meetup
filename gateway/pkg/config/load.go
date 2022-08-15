package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Load(path string) (*Config, error) {
	if err := readFromFile(path); err != nil {
		return nil, err
	}
	readFromEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func readFromFile(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func readFromEnv() {
	viper.SetEnvPrefix("gateway")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
