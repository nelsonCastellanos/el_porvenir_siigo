package controllers

import (
	"github.com/gin-gonic/gin"
)

// Structs

type HealthCheckController struct {
}

// Methods

func (ctrl HealthCheckController) HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Factory

func NewHealthCheckController() HealthCheckController {
	return HealthCheckController{}
}
