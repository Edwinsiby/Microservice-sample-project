package gateway

import (
	"gateway/pkg/middleware"

	"github.com/gin-gonic/gin"

	"gateway/pkg/handlers"
)

func NewAPIGateway() *gin.Engine {
	router := gin.Default()

	router.GET("/service1", handlers.HealthCheckService1)
	router.POST("/service1/signup", handlers.Signup)
	router.POST("/service1/login", handlers.Login)
	router.POST("/service1/addaddress", middleware.UserRetriveCookie, handlers.AddAddress)
	router.DELETE("/service1/logout", handlers.Logout)
	router.POST("/service1/adminsignup", handlers.AdminSignup)
	router.POST("/service1/adminlogin", handlers.AdminLogin)

	router.GET("/service2", handlers.HealthCheckService2)
	router.GET("/service2/productlist", middleware.UserRetriveCookie, handlers.ProductList)
	router.GET("/service2/productdetails/:id", middleware.UserRetriveCookie, handlers.ProductDetails)
	router.POST("/service2/addtocart/:id", middleware.UserRetriveCookie, handlers.AddToCart)
	router.DELETE("/service2/removefromcart/:id", middleware.UserRetriveCookie, handlers.RemoveFromCart)
	router.GET("/service2/cartdetails", middleware.UserRetriveCookie, handlers.CartDetails)

	router.GET("/service3", handlers.HealthCheckService3)
	router.POST("/service3/placeorder", middleware.UserRetriveCookie, handlers.PlaceOrder)
	router.POST("/service3/cancelorder", middleware.UserRetriveCookie, handlers.CancelOrder)
	router.GET("/service3/orderhistory", middleware.UserRetriveCookie, handlers.OrderHistory)

	router.GET("/service4", handlers.HealthCheckService4)
	router.POST("/service4/addproduct", middleware.AdminRetriveCookie, handlers.AddProduct)
	router.DELETE("/service4/removeproduct/:id", middleware.AdminRetriveCookie, handlers.RemoveProduct)

	router.GET("/gateway", handlers.HealthCheckGateway)

	return router
}
