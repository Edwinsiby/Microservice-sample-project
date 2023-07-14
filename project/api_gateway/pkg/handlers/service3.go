package handlers

import (
	"context"
	"fmt"
	"gateway/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var conn3 *grpc.ClientConn

func init() {
	var err error
	conn3, err = grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
}

func HealthCheckService3(c *gin.Context) {
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

func PlaceOrder(c *gin.Context) {
	fmt.Println("place order")
}

func CancelOrder(c *gin.Context) {
	fmt.Println("cancel order")
}

func OrderHistory(c *gin.Context) {
	fmt.Println("order history")
}
