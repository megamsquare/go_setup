package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	config "github.com/megamsquare/go_setup/pkg/env_config"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Address string `config:"GRPC_ADDRESS"`
	Port    int    `config:"GRPC_PORT" default:"9090"`
}

func Load_config() *GRPCConfig {
	var conf GRPCConfig
	config.Load(&conf)
	return &conf
}

func StartGRPCServer(conf GRPCConfig) {
	port := conf.Port
	if port == 0 {
		port = 9090
	}

	address := fmt.Sprintf("%s:%d", conf.Address, port)
	list, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("failed to listen to address: %s, with error: %v", address, err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Starting gRPC server on: %s", address)

	//Run gRPC server in goroutine to implement graceful shutdown.
	go func() {
		err := grpcServer.Serve(list)
		if err != nil {
			log.Printf("failed to serve gRPC server, with error: %v", err)
		}
	}()

	// Graceful Shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Block until signal is received
	<-signals

	log.Print("gRPC server shutting down!!!")
	grpcServer.GracefulStop()
}
