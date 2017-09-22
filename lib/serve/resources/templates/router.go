package templates

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

	templatesGroup := G.Group("/templates")

	// HOFSTADTER_START router-start
	// HOFSTADTER_END   router-start

	// names: serve | templates
	// routes NOT SAME NAME

	// routes RESOURCE
	templatesGroup.POST("/:templateId/render", Handle_POST_Render)

	// methods
	templatesGroup.GET("", Handle_LIST_Templates)
	templatesGroup.POST("", Handle_POST_Templates)
	templatesGroup.GET("/:templateId", Handle_GET_Templates)
	templatesGroup.PUT("/:templateId", Handle_PUT_Templates)
	templatesGroup.DELETE("/:templateId", Handle_DELETE_Templates)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

// HOFSTADTER_BELOW
