package handlers

import (
	"context"
	"gateway/pb"
	"gateway/pkg/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var conn4 *grpc.ClientConn

func init() {
	var err error
	conn4, err = grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
}

func HealthCheckService4(c *gin.Context) {
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

func AddProduct(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := int32(userID.(int))
	var input entity.Apparel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	client := pb.NewMyServiceClient(conn4)
	req := &pb.AddProductRequest{
		Name:     input.Name,
		Price:    int32(input.Price),
		ImageURL: input.ImageURL,
		Category: input.Category,
		AdminId:  int32(userId),
	}
	resp, err := client.AddProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.String(http.StatusOK, resp.Result)
	}
}

func RemoveProduct(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	}
	client := pb.NewMyServiceClient(conn4)
	req := &pb.RemoveProductRequest{
		Productid: int32(Id),
	}
	resp, err := client.RemoveProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Apparel not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}
