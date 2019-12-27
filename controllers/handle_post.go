package controllers

import (
	"encoding/json"
	"fmt"
	services "go_microservice/services/article"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// HandlePost parses request into services.Article struct
func HandlePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var a *services.Article
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(b, &a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", a)
	}
}
