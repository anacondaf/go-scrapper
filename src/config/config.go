package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/gorm/schema"
)

type database struct {
	ConnectionString string `mapstructure:"ConnectionString"`
	GormConfig       struct {
		NamingStrategy schema.NamingStrategy `mapstructure:"NamingStrategy"`
	} `mapstructure:"GormConfig"`
}

type cache struct {
	Address string `mapstructure:"Address"`
}

type rabbitmq struct {
	ConnectionString string `mapstructure:"ConnectionString"`
}

type sentry struct {
	DSN string `mapstructure:"DSN"`
}

type Config struct {
	Database database `mapstructure:"Database"`
	Cache    cache    `mapstructure:"Redis"`
	Rabbitmq rabbitmq `mapstructure:"Rabbitmq"`
	Sentry   sentry   `mapstructure:"Sentry"`
}

func LoadConfig(path string, logger *zerolog.Logger) (*Config, error) {
	viper.AddConfigPath(fmt.Sprintf("%v/src/config", path))

	// Read database.json
	viper.SetConfigName("database")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Read cache.json
	viper.SetConfigName("cache")
	err = viper.MergeInConfig()
	if err != nil {
		return nil, err
	}

	// Read rabbitmq.json
	viper.SetConfigName("rabbitmq")
	err = viper.MergeInConfig()
	if err != nil {
		return nil, err
	}

	// Read sentry.json
	viper.SetConfigName("sentry")
	err = viper.MergeInConfig()
	if err != nil {
		return nil, err
	}

	var config = Config{}

	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	logger.Info().Msg("Load Config Success")
	return &config, nil
}
