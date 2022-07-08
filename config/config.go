package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/phuslu/log"
	"github.com/rostikts/social_network/infrastructure/datastore"
)

type Config struct {
	SecretKey string `envconfig:"SECRET_KEY"`
	Database  datastore.Config
}

func NewConfig() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.DefaultLogger.Fatal().Err(err).Msg("Error occurred during config init")
	}
	return cfg
}
