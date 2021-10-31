package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DBUsername string `env:"DATABASE_USERNAME" env-default:"root"`
	DBHost     string `env:"DATABASE_HOST_NAME" env-default:"localhost"`
	DBPassword string `env:"DATABASE_PASSWORD" env-default:"root"`
	DBName     string `env:"DATABASE_NAME" env-default:"test"`
	DBSSLMode  string `env:"DATABASE_SSL_MODE" env-default:"disable"`
	HTTPAddr   string `env:"HTTP_PORT" env-default:"127.0.0.1:5000"`
}

var Cfg Config

func init() {
	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		panic(err)
	}
}
