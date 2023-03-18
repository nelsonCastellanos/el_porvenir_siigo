package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonCastellanos/golang-api-mongo/cmd/api/controllers"
)

type Routes struct {
	healthCheckController controllers.HealthCheckController
}

func main() {
	routes := Routes{
		healthCheckController: controllers.NewHealthCheckController(),
	}
	r := mapRoutes(routes)
	r.Run()
}

func mapRoutes(routes Routes) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", routes.healthCheckController.HandlePing)
	return r
}
