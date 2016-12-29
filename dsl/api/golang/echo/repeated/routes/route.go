{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package routes

import (
	"net/http"

	"github.com/labstack/echo"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
API:     {{API.name}}
Name:    {{RC.name}}
Route:   {{RC.route}}
Method:  {{RC.method}}
Path:    {{RC.path}}
Parent:  {{RC.parent}}
*/



// Should find a way to build up errors and return all
// {{upper RC.method}}  {{RC.input}}  ->  {{RC.output}}
func Handle_{{RC.method}}_{{RC.name}}(ctx echo.Context) error {
	// Check params
{{#params . as |P|}}

	{{#if (builtin P.type)}}
		// Extract
		{{camel P.name}} := ctx.QueryParam("{{P.name}}")
		// Validate
		{{camel P.name}}_tag := "required{{#each validation}},{{.}}{{/each}}"
		err := validate.Var({{camel P.name}}, {{camel P.name}}_tag)
		if err != nil {
			return err
		}
	{{else}}
			// Extract:
			// need to import the type and call pkg.New...
			input := new({{join2 "." (lower parent) (camelT name) }})
			if err := ctx.Bind(input); err != nil {
				return err
			}
			// Validate:
			// need to import the type and call pkg.New...
			if err := ctx.Validate(input); err != nil {
				return err
			}
	{{/if}}

{{/params}}


	// HOFSTADTER_START handler
	// HOFSTADTER_END   handler
}

{{/with}}
{{/with}}
