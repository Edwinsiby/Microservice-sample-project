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

var conn2 *grpc.ClientConn

func init() {
	var err error
	conn2, err = grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
}

func HealthCheckService2(c *gin.Context) {
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

func ProductList(c *gin.Context) {
	client := pb.NewMyServiceClient(conn2)
	req := &pb.Request{
		Data: "Mydata",
	}
	resp, err := client.ProductList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	responseList := make([]entity.Apparel, len(resp.Apparels))
	for i, apparel := range resp.Apparels {
		responseList[i] = entity.Apparel{
			ID:       int(apparel.Id),
			Name:     apparel.Name,
			Price:    int(apparel.Price),
			ImageURL: apparel.ImageURL,
			Category: apparel.Category,
		}
	}
	c.JSON(http.StatusOK, gin.H{"Apperals": responseList})
}

func ProductDetails(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
	}
	client := pb.NewMyServiceClient(conn2)
	req := &pb.ProductDetailsRequest{
		Productid: int32(Id),
	}
	resp, err := client.ProductDetails(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Apparel not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Apparel})
	}
}

func AddToCart(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	strId := c.Param("id")
	Id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	client := pb.NewMyServiceClient(conn2)
	req := &pb.AddToCartRequest{
		Userid:    int32(userId),
		Productid: int32(Id),
	}
	resp, err := client.AddToCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Add to cart failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

func RemoveFromCart(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "str conversion failed"})
		return
	}
	client := pb.NewMyServiceClient(conn2)
	req := &pb.AddToCartRequest{
		Userid:    int32(userId),
		Productid: int32(Id),
	}
	resp, err := client.RemoveFromCart(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Remove from cart failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp.Result})
	}
}

func CartDetails(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := userID.(int)
	client := pb.NewMyServiceClient(conn2)
	req := &pb.CartDetailsRequest{
		Userid: int32(userId),
	}
	resp, err := client.CartDetails(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Cart Not Found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": resp})
	}
}
