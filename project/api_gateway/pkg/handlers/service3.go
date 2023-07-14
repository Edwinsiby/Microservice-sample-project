package handlers

import (
	"context"
	"gateway/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func HealthCheckService3(c *gin.Context) {
	conn3, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
	defer conn3.Close()
	client := pb.NewMyServiceClient(conn3)
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := client.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}
	c.String(http.StatusOK, resp.Result)
}
