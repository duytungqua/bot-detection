package config

import (
	"github.com/go-playground/locales/vo"
	"github.com/spf13/viper"
)

type Config struct {
	Kafka struct {
		Brokers []string
		Topic   string
		GroupID string
	}
	API struct {
		Endpoint string
	}
}
var Cfg Config


func Load() Config{
	// Load configuration from a file or environment variables

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read configuration: " + err.Error())

	}
	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic("Failed to unmarshal configuration: " + err.Error())
	}
	
	Cfg.API.Endpoint = "https://api.example.com"
	return Cfg	

}