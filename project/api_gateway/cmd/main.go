package main

import (
	"log"
	"net/http"

	"gateway/pkg/gateway"
)

func main() {
	apiGateway := gateway.NewAPIGateway()

	log.Println("API Gateway is running on :8080")
	if err := http.ListenAndServe(":8080", apiGateway); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
