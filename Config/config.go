package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	RabbitMQ   string
	RedisHost  string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config") 
	viper.SetConfigType("yaml")   // File format
	viper.AddConfigPath("configuration/config.go")      

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error in reading config file: %s", err))
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode config: %s", err))
	}
}
