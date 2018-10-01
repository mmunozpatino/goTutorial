package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", hello)

	router.POST("/user",postHello)

	router.Run(":8080")
}
