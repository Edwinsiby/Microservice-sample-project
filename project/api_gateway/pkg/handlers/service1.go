package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"gateway/pb"
)

func HealthCheckService1(c *gin.Context) {
	conn2, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
	defer conn2.Close()
	client := pb.NewMyServiceClient(conn2)
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := client.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}
	c.String(http.StatusOK, resp.Result)
}

func Signup(c *gin.Context) {

	conn1, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
	defer conn1.Close()
	client := pb.NewMyServiceClient(conn1)
	req := &pb.SignupRequest{
		Firstname: "Edwin",
		Lastname:  "Siby",
		Email:     "edwin@gmail.com",
		Phone:     "9048402133",
		Password:  "pass@123",
	}
	resp, err := client.Signup(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}
	c.String(http.StatusOK, resp.Result)
}

func Login(c *gin.Context) {
	fmt.Println("hi")
}

func AddAddress(c *gin.Context) {
	fmt.Println("hi")
}
