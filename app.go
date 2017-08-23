package main

import (
	// Gin framework and templator
	"github.com/gin-gonic/gin"
	"github.com/madhums/go-gin-mgo-demo/gin_html_render"
	// Route handlers
	"./app/routes/welcome"
	"./app/routes/users"
)

func main() {
	app := gin.Default()

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "layouts/application"
	htmlRender.TemplatesDir = "app/views/"

	// Tell gin to use our html render
	app.HTMLRender = htmlRender.Create()

	// Static Public Files
	app.Static("/stylesheets", "./public/stylesheets")

	// Upload routes
	welcome.Routes(app)
	users.Routes(app)

	// Run app
	app.Run(":3000")
}
