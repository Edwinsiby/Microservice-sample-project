package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckGateway(c *gin.Context) {

	c.String(http.StatusOK, "This is the GET endpoint")

}
