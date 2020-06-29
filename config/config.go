package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"server_port"`

	DatabaseHost     string `envconfig:"database_host"`
	DatabaseUsername string `envconfig:"database_username"`
	DatabasePassword string `envconfig:"database_password"`
	DatabaseName     string `envconfig:"database_name"`

	ShutdownPeriod int `envconfig:"shutdown_period"`

	NsqHost string `envconfig:"nsq_host"`
	NsqPort string `envconfig:"nsq_port"`
}

var conf Config

func Get() *Config {
	once := sync.Once{}
	once.Do(func() {
		err := envconfig.Process("", &conf)
		if err != nil {
			log.Fatalf("err conf : %v", err)
		}
	})

	return &conf
}
