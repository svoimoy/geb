{{#with DslContext as |API| }}
package main

import (
	"github.com/labstack/echo"

	{{#if DslContext.resources}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources"
	{{/if}}
	{{#if DslContext.routes}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/routes"
	{{/if}}

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

	// Routes
	{{#each routes as |R|}}
	{{> router/route.go R ~}}
	{{/each}}

	// Resources
	{{#each resources as |R| ~}}
		{{#each methods as |M| ~}}
	{{> router/resource.go R ~}}
		{{/each}}
	{{/each }}

	return nil
}
/*
	This is a test for dotpath indexing enhancement:

	Path: DslContext.[???]

	??? = routes.version
	---
	{{#get_elem_by_name "routes.version" "" true data=API}}
	{{name}}
	{{/get_elem_by_name}}
	---
	{{#dotpath "api.routes.version" @root.DslContext true}}
	{{name}}
	{{/dotpath}}
	---
	{{#dotpath "routes.version" API true}}
	{{{.}}}
	{{/dotpath}}
  ---
	{{#getdsl "api.routes.[0]" true}}
	{{#each .}}
	 - {{name}}
	{{/each~}}
	{{/getdsl}}
  ---
	{{#gettype "count" true}}
	{{name}}
	{{/gettype}}

*/

{{/with}}
