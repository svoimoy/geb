package routes

import (
	"github.com/labstack/echo"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

func InitRouter(G *echo.Group) (err error) {

	// HOFSTADTER_START router-pre
	// HOFSTADTER_END   router-pre

	serveGroup := G.Group("")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: serve | serve
	// routes SAME NAME

	addPrometheusHandlers(serveGroup)

	addKubernetesHandlers(serveGroup)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
