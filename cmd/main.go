package main

import (
	"log"

	"github.com/EgoitzAB/ecom_api_exercise"
)
func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}