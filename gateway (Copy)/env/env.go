package env

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ENV                       string
	SERVER_PORT               string
	SERVER_HOST               string
	SERVER_HOST_BE            string
	SERVER_HOST_FE            string
	SERVER_TIMEZONE           string
	SERVICE_EXAMPLE_PORT      string
	SERVICE_AUTH_PORT         string
}

var cfg Config

func LoadEnvironmentFile() {
	viper.SetConfigFile("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("environment: %v", err.Error())
		os.Exit(1)
	}
	fmt.Println("environment: load successfully")
}

func NewEnv() *Config {
	if cfg.ENV == "" {
		cfg = Config{
			ENV:                       viper.GetString("env"),
			SERVER_HOST:               viper.GetString("server.host"),
			SERVER_PORT:               viper.GetString("server.port"),
			SERVER_HOST_BE:            viper.GetString("server.host_be"),
			SERVER_HOST_FE:            viper.GetString("server.host_fe"),
			SERVER_TIMEZONE:           viper.GetString("server.timezone"),
			SERVICE_EXAMPLE_PORT:      viper.GetString("service.example.port"),
			SERVICE_AUTH_PORT:         viper.GetString("service.auth.port"),
			// Initialize other configuration variables here
		}
	}
	return &cfg
}
