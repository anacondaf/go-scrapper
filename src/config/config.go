package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm/schema"
)

type database struct {
	ConnectionString string `mapstructure:"ConnectionString"`
	GormConfig       struct {
		NamingStrategy schema.NamingStrategy `mapstructure:"NamingStrategy"`
	} `mapstructure:"GormConfig"`
}

type Config struct {
	database `mapstructure:",squash"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(fmt.Sprintf("%v/src/config", path))
	viper.SetConfigName("database")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config = Config{}

	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	fmt.Println("Load config success!")
	return &config, nil
}
