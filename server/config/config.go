package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string   `default:"8001"`
	Env  Platform `default:"local"`
	Database
}

type Platform string

const (
	Local Platform = "local"
	Dev   Platform = "dev"
	Stg   Platform = "stg"
	Prod  Platform = "prod"
)

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER"  default:"postgres"`
	Password string `envconfig:"DB_PASSWORD"  default:"password"`
	Name     string `envconfig:"DB_NAME"  default:"postgres"`
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
