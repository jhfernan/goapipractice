package main

import (
	"github.com/gin-gonic/gin"
	// Route handlers
	"./app/routes/welcome"
	"./app/routes/users"
)

func main() {
	app := gin.Default()
	app.Delims("<<", ">>")
	app.LoadHTMLGlob("app/views/**/*")
	// Upload routes
	welcome.Routes(app)
	users.Routes(app)

	// Run app
	app.Run(":3000")
}
