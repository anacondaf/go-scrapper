package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

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
