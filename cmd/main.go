package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/megamsquare/go_setup/config"
	database "github.com/megamsquare/go_setup/db"
)

func main() {
	err := config.Set_env(".env")
	if err != nil {
		log.Println("env file load config error: ", err)
	}
	http.HandleFunc("/", handler)
	log.Println("Starting Server at port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Starting Server")
}
