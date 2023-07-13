package service

import (
	"context"
	"log"

	pb "service2/pb"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyService) MyMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Microservice2: MyMethod called")

	// Implement the logic for MyMethod
	result := "Hola, " + req.Data
	return &pb.Response{Result: result}, nil
}
