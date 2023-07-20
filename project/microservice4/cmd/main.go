package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "service4/pb"
	"service4/pkg/service"
)

func main() {
	myService := &service.MyService{}

	grpcServer := grpc.NewServer()

	pb.RegisterMyServiceServer(grpcServer, myService)

	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Microservice4 is running...")
	go grpcServer.Serve(listener)

	if err := http.ListenAndServe(":8084", nil); err != nil {
		log.Fatalf("Failed to start health check server: %v", err)
	}

}
