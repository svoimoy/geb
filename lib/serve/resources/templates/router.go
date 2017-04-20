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
	templatesGroup.POST("/:template-id/render", Handle_POST_Render)

	// methods
	templatesGroup.GET("", Handle_LIST_Templates)
	templatesGroup.POST("", Handle_POST_Templates)
	templatesGroup.GET("/:template-id", Handle_GET_Templates)
	templatesGroup.PUT("/:template-id", Handle_PUT_Templates)
	templatesGroup.DELETE("/:template-id", Handle_DELETE_Templates)

	// HOFSTADTER_START router-end
	// HOFSTADTER_END   router-end

	return nil
}

/*

null


---

ctx_path: dsl.lib.serve.api.resources.[0]
methods:
- method: list
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[0].output.[0]
    name: ts
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[0]/output
    pkgPath: serve/templates/ts
    type: array:lib.templates.template.views.short
- input:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[1].input.[0]
    name: in-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[1]/input
    pkgPath: serve/templates/in-tpl
    type: lib.templates.template.views.create
  method: post
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[1].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[1]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template
- method: get
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[2].output.[0]
    name: t
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[2]/output
    pkgPath: serve/templates/t
    type: lib.templates.template
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[2].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[2]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
- input:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].input.[0]
    name: in-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/input
    pkgPath: serve/templates/in-tpl
    type: lib.templates.template
  method: put
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
- method: delete
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[4].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[4]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template.views.short
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[4].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[4]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
name: templates
omit-db-calls: true
parent: serve
parent_path: dsl.lib.serve.api
path: resources
pkg_path: lib/serve/api/resources
pkgPath: serve/templates
resource: lib.templates.template
route: templates
routes:
- ctx_path: dsl.lib.serve.api.resources.[0].routes.[0]
  method: POST
  name: render
  parent: serve.templates
  parent_path: dsl.lib.serve.api.resources.[0]
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].routes.[0].path-params.[0]
    name: template-id
    parent: serve.templates.render
    parent_path: dsl.lib.serve.api.resources.[0].routes.[0]
    pkg_path: lib/serve/api/resources/[0]/routes/[0]/path-params
    pkgPath: serve/templates/render/template-id
    type: lib.templates.template.fields.id
  pkg_path: lib/serve/api/resources/[0]/routes
  pkgPath: serve/templates/render
  route: render


*/

// HOFSTADTER_BELOW