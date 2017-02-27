{{#with DslContext as |API| }}
package main

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

/*
{{{yaml .}}}
*/
