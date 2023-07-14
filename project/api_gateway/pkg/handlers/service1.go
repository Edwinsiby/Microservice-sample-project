package handlers

import (
	"context"
	"log"
	"net/http"

	"gateway/pkg/entity"
	"gateway/pkg/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"gateway/pb"
)

var conn1 *grpc.ClientConn
var err error

func init() {
	conn1, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to service1: %v", err)
	}
}
func HealthCheckService1(c *gin.Context) {

	client := pb.NewMyServiceClient(conn1)
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
	var userInput entity.Signup
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := pb.NewMyServiceClient(conn1)
	req := &pb.SignupRequest{
		Firstname: userInput.FirstName,
		Lastname:  userInput.LastName,
		Email:     userInput.Email,
		Phone:     userInput.Phone,
		Password:  userInput.Password,
	}
	resp, err := client.Signup(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.String(http.StatusOK, resp.Result)
	}

}

func Login(c *gin.Context) {
	var userInput entity.Login
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := pb.NewMyServiceClient(conn1)
	req := &pb.LoginRequest{
		Phone:    userInput.Phone,
		Password: userInput.Password,
	}
	resp, err := client.Login(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		middleware.CreateJwtCookie(int(resp.Userid), userInput.Phone, "user", c)
		c.String(http.StatusOK, resp.Result)
	}
}

func AddAddress(c *gin.Context) {
	userID, _ := c.Get("userID")
	userId := int32(userID.(int))
	var address entity.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := pb.NewMyServiceClient(conn1)
	req := &pb.AddAddressRequest{
		Userid:  userId,
		House:   address.House,
		Street:  address.Street,
		City:    address.City,
		Pincode: address.Pincode,
		Type:    address.Type,
	}
	resp, err := client.AddAddress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"massage": resp.Result})
	}
}

func Logout(c *gin.Context) {
	err := middleware.DeleteCookie(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cookie deletion failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
	}
}

func AdminSignup(c *gin.Context) {
	var userInput entity.Admin
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := pb.NewMyServiceClient(conn1)
	req := &pb.AdminSignupRequest{
		Adminname: userInput.AdminName,
		Email:     userInput.Email,
		Phone:     userInput.Phone,
		Password:  userInput.Password,
		Role:      userInput.Role,
	}
	resp, err := client.AdminSignup(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.String(http.StatusOK, resp.Result)
	}

}

func AdminLogin(c *gin.Context) {
	var userInput entity.Login
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := pb.NewMyServiceClient(conn1)
	req := &pb.LoginRequest{
		Phone:    userInput.Phone,
		Password: userInput.Password,
	}
	resp, err := client.AdminLogin(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		middleware.CreateJwtCookie(int(resp.Userid), userInput.Phone, "admin", c)
		c.String(http.StatusOK, resp.Result)
	}
}
