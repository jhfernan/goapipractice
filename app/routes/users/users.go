package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	users := routes.Group("/users")

	// Route that returns all users
	users.GET("", func(c *gin.Context) {
		c.HTML(200, "users/users", gin.H{
			"title": "Go Testing | Users",
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
