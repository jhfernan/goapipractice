package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	users := routes.Group("/users")

	// Route that returns all users
	users.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": "all",
		})
	})
	// return name in JSON
	users.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
}
