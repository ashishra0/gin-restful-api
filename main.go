package main

import (
	"go_microservice/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", handler.HandlePost())
	r.Run(":8000")
}
