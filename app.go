package main

import (
	"github.com/gin-gonic/gin"
	// Route handlers
	"./routes/welcome"
	"./routes/users"
)

func main() {
	app := gin.Default()
	app.Delims("<<", ">>")
	app.LoadHTMLGlob("views/**/*")
	// Upload routes
	welcome.Routes(app)
	users.Routes(app)

	// Run app
	app.Run(":3000")
}
