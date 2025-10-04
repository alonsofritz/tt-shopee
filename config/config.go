package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName    string
	AppVersion string
	ServerPort string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	configFile := fmt.Sprintf(".env.%s", env)

	viper.SetConfigFile(configFile)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("APP_NAME", "tt-shopee")
	viper.SetDefault("APP_VERSION", "0.0.0")
	viper.SetDefault("SERVER_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Não foi possível carregar %s.", configFile)
	}

	return &Config{
		AppName:    viper.GetString("APP_NAME"),
		AppVersion: viper.GetString("APP_NAME"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}
}
