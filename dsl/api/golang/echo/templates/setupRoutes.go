{{#with DslContext as |API| }}
package main

import (
	"github.com/labstack/echo"

	{{#if API.resources}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources"
		{{#each API.resources as |R1| ~}}
			{{#if R1.resources}}
			"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources/{{R1.name}}"
			{{#each R1.resources as |R2| ~}}
				{{#if R2.resources}}
				"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/resources/{{R1.name}}/{{R2.name}}"
				{{/if}}
			{{/each}}	
			{{/if}}
		{{/each}}	
	{{/if}}

	{{#if API.routes}}
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
	{{#each API.resources as |R1| ~}}
		{{camel R1.route}}Group := G.Group("/{{R1.route}}")
		{{#each R1.methods as |M1| ~}}
			{{> router/resource.go R=R1 M=M1 ~}}
		{{/each}}

		{{#each R1.resources as |R2| ~}}
			{{camel R2.route}}Group := {{camel R1.route}}Group.Group("{{#each R2.path-params as |P|~}}/:{{P.name }}{{/each}}/{{R2.route}}")
			{{#each R2.methods as |M2| ~}}
				{{> router/resource.go R=R2 M=M2 ~}}
			{{/each}}

			{{#each R2.resources as |R3| ~}}
			{{camel R3.route}}Group := {{camel R2.route}}Group.Group("{{#each R3.path-params as |P|~}}/:{{P.name }}{{/each}}/{{R3.route}}")
				{{#each R3.methods as |M3| ~}}
					{{> router/resource.go R=R3 M=M3 ~}}
				{{/each}}

			{{/each }}
		{{/each }}
	{{/each }}

	return nil
}

{{/with}}
