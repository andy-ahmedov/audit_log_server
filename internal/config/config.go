package config

import "github.com/kelseyhightower/envconfig"

// создаем структуру Монго в котором будут 4 поля: uri username пароль и database
// создаем структуру Сервер с единственным полем порт
// создаем структуру конфиг с двумя структурами выше
//
// создаем функцию New()
// инициализируем структуру конфиг с помощью функции new
// используем библу енвконфиг и метод процесс для связывания данных из енв со структурами
// возвращаем структуру

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
