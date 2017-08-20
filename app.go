package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"greeting": "Hello World",
		})
	})

	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
	r.Run()
}
