package routes

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START start
// HOFSTADTER_END   start

func addPrometheusHandlers(G *echo.Group) (err error) {

	group := G.Group("")

	group.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return nil
}

// HOFSTADTER_BELOW
