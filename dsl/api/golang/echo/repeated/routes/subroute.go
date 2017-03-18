{{#with RepeatedContext as |RC| }}
{{#with DslContext as |API| }}
package routes
// package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
API:     {{API.name}}
Name:    {{RC.name}}
Route:   {{RC.route}}
Method:  {{RC.method}}
Path:    {{RC.path}}
Parent:  {{RC.parent}}

Params:
{{#params}}
{{name}} {{type}}  {{#if required}}(required){{/if}}
  {{#each validation}}
	- {{@key}}:  {{.}}
	{{/each}}
{{/params}}

*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init



// Should find a way to build up errors and return all
func {{RC.name}}_{{RC.method}}_Handler(ctx *gin.Context) {
	// Check params
{{#each params}}
	{{name}} := c.Query("{{name}}")
	{{#if required }}
	if {{name}} == "" {
		res := gin.H{"error": "missing {{name}} in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	{{/if}}
	{{#if type }} {{> (concat3 "parse/" type ".go") }} {{/if}}
{{/each}}

}

{{/with}}
{{/with}}
