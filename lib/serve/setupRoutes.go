package serve

import (
	"github.com/labstack/echo"

	"github.ibm.com/hofstadter-io/geb/lib/serve/resources"

	"github.ibm.com/hofstadter-io/geb/lib/serve/routes"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

func setupRoutes(G *echo.Group) error {

	// HOFSTADTER_START pre-routes
	// HOFSTADTER_END   pre-routes

	// Routes
	G.GET("/readyz", routes.Handle_GET_ReadyCheck)
	G.GET("/healthz", routes.Handle_GET_HealthCheck)

	// HOFSTADTER_START post-routes
	// HOFSTADTER_END   post-routes

	// HOFSTADTER_START pre-resources
	// HOFSTADTER_END   pre-resources

	// Resources
	templatesGroup := G.Group("/templates")
	templatesGroup.GET("", resources.Handle_LIST_Templates)
	templatesGroup.POST("", resources.Handle_POST_Templates)
	templatesGroup.GET("/:template-id", resources.Handle_GET_Templates)
	templatesGroup.PUT("/:template-id", resources.Handle_PUT_Templates)
	templatesGroup.DELETE("/:template-id", resources.Handle_DELETE_Templates)

	// HOFSTADTER_START post-resources
	// HOFSTADTER_END   post-resources

	return nil
}

// HOFSTADTER_BELOW
