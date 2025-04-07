package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	HTTP HttpConfig
	DB   DBConfig
}

type HttpConfig struct {
	Addr                  string `env:"HTTP_CONFIG_ADDR" default:":8080"`
	RequestTimeoutSeconds int    `env:"HTTP_CONFIG__REQUEST_TIMEOUT_SECONDS"   default:"60"`
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST" default:"localhost"`
	Port     string `env:"POSTGRES_PORT" default:"5432"`
	User     string `env:"POSTGRES_USER" default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" default:"postgres"`
	Database string `env:"POSTGRES_DB" default:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" default:"disable"`
}

func New(filenames ...string) (*Config, error) {
	cfg := new(Config)
	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	defaults.SetDefaults(cfg)

	return cfg, nil
}
