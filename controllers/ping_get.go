package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet endpoint handler function
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "World",
		})
	}
}
