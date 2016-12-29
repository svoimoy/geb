{{#with dsl.api as |API| }}
package main

/*
	This is a test for dotpath indexing enhancement:

	Path: dsl.api.[???]

	??? = routes.version
	---
	{{#get_elem_by_name "routes.version" "" data=API}}
	{{name}}
	{{/get_elem_by_name}}
	---
	{{#dotpath "api.routes.version" @root.dsl}}
	{{name}}
	{{/dotpath}}
	---
	{{#dotpath "routes.version" API}}
	{{name}}
	{{/dotpath}}
  ---
	{{#getdsl "api.routes.[0]"}}
	{{#each .}}
	 - {{name}}
	{{/each~}}
	{{/getdsl}}
  ---
	{{#gettype "count"}}
	{{name}}
	{{/gettype}}

*/

{{/with}}

