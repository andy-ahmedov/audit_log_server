package config

import "github.com/kelseyhightower/envconfig"

type Mongo struct {
	Uri      string
	Username string
	Password string
	Database string
}

type Server struct {
	Port int
}

type Config struct {
	Db     Mongo
	Server Server
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("db", &cfg.Db); err != nil {
		return nil, err
	}
	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	return cfg, nil
}
