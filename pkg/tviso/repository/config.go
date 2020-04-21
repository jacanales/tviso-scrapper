package repository

import "github.com/kelseyhightower/envconfig"

type Config struct {
	APIAddr string `envconfig:"ENDPOINT" default:"localhost:8080"`
	Cookie string `envconfig:"COOKIE"`
}

func NewConfig() Config {
	cfg := Config{}

	envconfig.MustProcess("TVISO", &cfg)

	return cfg
}
