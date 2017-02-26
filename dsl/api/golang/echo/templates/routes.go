{{#with DslContext as |API| }}
package main

/*
	This is a test for dotpath indexing enhancement:

	Path: dsl.api.[???]

	??? = routes.version
	---
	{{#get_elem_by_name "routes.version" "" true data=API}}
	{{name}}
	{{/get_elem_by_name}}
	---
	{{#dotpath "api.routes.version" @root.dsl true}}
	{{name}}
	{{/dotpath}}
	---
	{{#dotpath "routes.version" API true}}
	{{name}}
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

