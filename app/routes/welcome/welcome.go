package welcome

import (
	"github.com/gin-gonic/gin"
)


func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.HTML(200, "welcome.html", gin.H{ "title": "Home" })
	})
}
