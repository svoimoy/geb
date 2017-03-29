{{#with RepeatedContext as |CTX| }}
{{#with DslContext as |API| }}
package resources
// package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	"net/http"

	"github.com/labstack/echo"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
API:       {{API.name}}
Name:      {{CTX.name}}
Route:     {{CTX.route}}
Resource:  {{CTX.resource}}
Path:      {{CTX.path}}
Parent:    {{CTX.parent}}

*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init



{{#each methods}}

{{#with . as |M|}}
// {{upper M.method}}  {{M.input}}  ->  {{M.output}}
func Handle_{{upper M.method}}_{{camelT CTX.name}}(ctx echo.Context) error {

	// input
	{{#if M.path-params}}
	// path param
		{{> api/golang/echo/input/path-params.go PARAMS=M.path-params }}
	{{/if}}

	{{#if M.input}}
	// START binding input to query/form/body params
	// Initialize
	{{> type/golang/var-new.go typename=M.input }}

	// END binding input to query/form/body params
	{{/if}}

	{{#if M.params}}
	// START indep query/form params

	{{#params . as |P|}}
	{{#if (builtin P.type)}}
		// Extract {{P.name}}
		{{camel P.name}} := ctx.QueryParam("{{P.name}}")

		// Validate {{p.name}} 
		{{camel P.name}}_tag := "required{{#each validation}},{{.}}{{/each}}"
		err := validator.New().Var({{camel P.name}}, {{camel P.name}}_tag)
		if err != nil {
			return err
		}
	{{else}}
	   // Only built in types are supported in query/form params. Use 'input' option on the resource.method
	{{/if}}

	{{/params}}
	// END query/form/body params
	{{/if}}


{{#if M.output}}
	// OUTPUT
	{{> type/golang/var-new.go typename=M.output }}
{{else}}
	// NO OUTPUT
{{/if}}

	// HOFSTADTER_START {{lower M.method}}
	// HOFSTADTER_END   {{lower M.method}}

	// return the output response

	return ctx.JSON(http.StatusOK, output)

}
{{/with}}

{{/each}} // End resource.methods

/*
{{{yaml CTX}}}
*/

{{/with}}
{{/with}}

// HOFSTADTER_BELOW
