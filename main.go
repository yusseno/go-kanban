package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test message",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
