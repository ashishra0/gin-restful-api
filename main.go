package main

import (
	"go_microservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", controllers.HandlePost())
	r.Run(":8000")
}
