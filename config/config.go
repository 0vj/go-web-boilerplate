package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUsername string `envconfig:"DATABASE_USERNAME" default:"root"`
	DBHost     string `envconfig:"DATABASE_HOST" default:"localhost"`
	DBPassword string `envconfig:"DATABAE_PASSWORD" default:"root"`
	DBName     string `envconfig:"DATABASE_NAME" default:"schotori"`
	DBSSLMode  string `envconfig:"DATABSE_SSLMODE" default:"disable"`
	HTTPAddr   string `envconfig:"HTTP_PORT" default:"127.0.0.1:5000"`
}

var Cfg Config

func init() {
	godotenv.Load()
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}
