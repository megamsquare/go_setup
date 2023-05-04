package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	config "github.com/megamsquare/go_setup/pkg/env_config"
)

type ServerConfig struct {
	Address string `config:"HTTP_ADDRESS"`
	Port    int    `config:"HTTP_PORT" default:"3000"`
	ReadTimeout time.Duration `config:"HTTP_READ_TIMEOUT" default:"10s"`
	WriteTimeout      time.Duration `config:"HTTP_WRITE_TIMEOUT" default:"120s"`
	ShutdownTimeout   time.Duration `config:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
}

func Load_config() *ServerConfig {
	var conf ServerConfig
	config.Load(&conf)
	return &conf
}

func ListenAndServe(conf ServerConfig, handler http.Handler) {
	port := conf.Port
	if port == 0 {
		appEnginePort := os.Getenv("PORT")
		if appEnginePort == "" {
			port = 3000
		}
	}

	address := fmt.Sprintf("%s:%d", conf.Address, port)

	srv := http.Server{
		Addr:         address,
		Handler: handler,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	}

	log.Printf("Server is listening on: %s", address)

	//Run server in goroutine to implement graceful shutdown.
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	// Graceful Shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Block until signal is received
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
	defer cancel()
	log.Println("Server is shutting down!!!")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Server shutting down with error: %v\n", err)
	}
}

