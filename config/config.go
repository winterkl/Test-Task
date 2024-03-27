package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Pg PSQL `yaml:"postgres"`
	}

	PSQL struct {
		Host     string `env-required:"true" yaml:"host"`
		Port     string `env-required:"true" yaml:"port"`
		User     string `env-required:"true" yaml:"user"`
		Password string `env-required:"true" yaml:"password"`
		DbName   string `env-required:"true" yaml:"db_name"`
		SslMode  string `env-required:"true" yaml:"ssl_mode"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
