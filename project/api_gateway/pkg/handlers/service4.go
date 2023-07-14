package handlers

import (
	"context"
	"gateway/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func HealthCheckService4(c *gin.Context) {
	conn4, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
	defer conn4.Close()
	client := pb.NewMyServiceClient(conn4)
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := client.MyMethod(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call MyMethod: %v", err)
	}
	c.String(http.StatusOK, resp.Result)
}
