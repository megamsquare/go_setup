package server

import (
	"time"

	config "github.com/megamsquare/go_setup/pkg/env_config"
)

type ServerConfig struct {
	Address string `config:"HTTP_ADDRESS"`
	Port    int    `config:"HTTP_PORT" default:"3000"`
	ReadTimeout time.Duration `config:"HTTP_READ_TIMEOUT" default:"10s"`
	ShutdownTimeout   time.Duration `config:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
}

func Load_config() *ServerConfig {
	var conf ServerConfig
	config.Load(&conf)
	return &conf
}

