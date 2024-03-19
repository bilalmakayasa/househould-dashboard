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

func DbConfig() string {
	return fmt.Sprintf("sslmode=disable dbname=%s host=%s port=%s user=%s password=%s", viper.GetString("database.name"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.user"), viper.GetString("database.password"))
}

func GetJwtKey() string {
	return viper.GetString("jwt.secret")
}
