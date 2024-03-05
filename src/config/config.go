package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigName("default")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./src/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
