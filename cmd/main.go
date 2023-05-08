package main

import (
	"log"

	database "github.com/megamsquare/go_setup/pkg/db"
	config "github.com/megamsquare/go_setup/pkg/env_config"
	"github.com/megamsquare/go_setup/pkg/grpc"
	"github.com/megamsquare/go_setup/pkg/queue"
)

func main() {
	// Call enveronment var
	err := config.Set_env(".env")
	if err != nil {
		log.Println("env file load config error: ", err)
	}

	// Init RabbitMQ
	rmqConfig := queue.Load_config()
	rabbitMQ, err := queue.Connect_rabbitMQ(rmqConfig)
	if err != nil {
		log.Printf("Error connectiong to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Init DB
	dbConfig := database.Load_config()
	db, err := database.Connect_db(dbConfig)
	if err != nil {
		log.Printf("Database error: %v", err)
	}
	defer db.Close()

	gRPCconfig := grpc.Load_config()
	grpc.StartGRPCServer(*gRPCconfig)

	// http.HandleFunc("/", handler)
	// log.Println("Starting Server at port :3000")
	// log.Fatal(http.ListenAndServe(":3000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Starting Server")
// }
