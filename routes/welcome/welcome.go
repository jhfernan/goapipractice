package welcome

import (
	"github.com/gin-gonic/gin"
)


func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.HTML(200, "welcome/home.html", gin.H{
			"title": "Hello World",
		})
	})

	route.GET("/lol", func(c *gin.Context) {
		c.HTML(200, "home/home.html", gin.H{
			"title": "Hello World",
		})
	})
}
