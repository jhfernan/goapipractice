package welcome

import (
	"github.com/gin-gonic/gin"
)


func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.HTML(200, "home/index.tmpl", gin.H{
			"title": "Hello World",
		})
	})

	route.GET("/lol", func(c *gin.Context) {
		c.HTML(200, "welcome/index.tmpl", gin.H{
			"title": "Hello World",
		})
	})
}