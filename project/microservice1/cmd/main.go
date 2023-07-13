package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "service1/pb"
	"service1/pkg/service"
)

func main() {
	myService := &service.MyService{}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the service with the server
	pb.RegisterMyServiceServer(grpcServer, myService)

	// Start the server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Microservice1 is running...")
	go grpcServer.Serve(listener)

	// Start the HTTP server for health checks
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}

}
