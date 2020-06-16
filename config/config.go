package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"server_port"`

	DatabaseHost     string `mapstructure:"database_host"`
	DatabaseUsername string `mapstructure:"database_username"`
	DatabasePassword string `mapstructure:"database_password"`
	DatabaseName     string `mapstructure:"database_name"`

	NsqHost string `mapstructure:"nsq_host"`
	NsqPort string `mapstructure:"nsq_port"`
}

var conf Config

func Get() *Config {
	viper.SetConfigName("dev.config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalln(err)
	}
	return &conf
}
