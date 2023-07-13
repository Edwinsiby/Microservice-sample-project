package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "service2/pb"
	"service2/pkg/service"
)

func main() {
	// Create a gRPC server
	server := grpc.NewServer()

	// Initialize the service implementation
	myService := &service.MyService{}

	// Register the service with the server
	pb.RegisterMyServiceServer(server, myService)

	// Start the server
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Microservice2 is running...")
	go server.Serve(listener)

	// Start the HTTP server for health checks
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}
}
