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

	server := grpc.NewServer()

	myService := &service.MyService{}

	pb.RegisterMyServiceServer(server, myService)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Microservice2 is running...")
	go server.Serve(listener)

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}
}
