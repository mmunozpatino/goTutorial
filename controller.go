package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Hello %s!", name),
	})
}

func postHello(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
