package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type HTTP struct {
	Port string `env:"HTTP_PORT" env-default:"8080"`
}

type DB struct {
	Name     string `env:"DB_NAME" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-default:"postgres"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

type Config struct {
	HTTP
	DB
}

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "reading config error")
	}
	return &cfg, nil
}
