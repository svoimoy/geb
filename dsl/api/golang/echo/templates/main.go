{{#with dsl.api}}
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources"
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/routes"

)

// Name:     {{name}}
// Version:  {{version}}
// About:    {{about}}

func main() {

	E := echo.New()

	// Pre-Middleware

	// Use-Middleware
	E.Use(middleware.Logger())
	E.Use(middleware.Recover())

	// Routes
	G := E.Group("{{config.base-url}}")
	{{#each routes}}
	{{> router/route.go .}}
	{{/each}}

	// Resources
	{{#each resources}}
	{{> router/resource.go .}}
	{{/each}}

	E.Logger.SetLevel(log.INFO)
	E.Logger.Fatal(E.Start("localhost:1323"))
}
{{/with}}
