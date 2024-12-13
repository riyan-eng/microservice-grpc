package env

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ENV                 string
	SERVER_PORT         string
	SERVER_HOST         string
	SERVER_HOST_BE      string
	SERVER_HOST_FE      string
	SERVER_TIMEZONE     string
	JWT_SECRET_ACCESS   string
	JWT_SECRET_REFRESH  string
	JWT_SECRET_RESET    string
	JWT_EXPIRED_LOGOFF  time.Duration
	JWT_EXPIRED_ACCESS  time.Duration
	JWT_EXPIRED_REFRESH time.Duration
	JWT_EXPIRED_RESET   time.Duration
	SMTP_HOST           string
	SMTP_PORT           string
	SMTP_EMAIL          string
	SMTP_PASSWORD       string
	POSTGRE_HOST        string
	POSTGRE_PORT        string
	POSTGRE_TIMEZONE    string
	POSTGRE_USERNAME    string
	POSTGRE_PASSWORD    string
	POSTGRE_DATABASE    string
	REDIS_HOST          string
	REDIS_PORT          string
	REDIS_USERNAME      string
	REDIS_PASSWORD      string
	REDIS_DATABASE      int
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
			ENV:              viper.GetString("env"),
			SMTP_HOST:        viper.GetString("smtp.host"),
			SMTP_PORT:        viper.GetString("smtp.port"),
			SMTP_EMAIL:       viper.GetString("smtp.email"),
			SMTP_PASSWORD:    viper.GetString("smtp.password"),
			POSTGRE_HOST:     viper.GetString("postgre.host"),
			POSTGRE_PORT:     viper.GetString("postgre.port"),
			POSTGRE_TIMEZONE: viper.GetString("postgre.timezone"),
			POSTGRE_USERNAME: viper.GetString("postgre.username"),
			POSTGRE_PASSWORD: viper.GetString("postgre.password"),
			POSTGRE_DATABASE: viper.GetString("postgre.database"),
			REDIS_HOST:       viper.GetString("redis.host"),
			REDIS_PORT:       viper.GetString("redis.port"),
			REDIS_USERNAME:   viper.GetString("redis.username"),
			REDIS_PASSWORD:   viper.GetString("redis.password"),
			REDIS_DATABASE:   viper.GetInt("redis.database"),
			// Initialize other configuration variables here
		}
	}
	return &cfg
}
