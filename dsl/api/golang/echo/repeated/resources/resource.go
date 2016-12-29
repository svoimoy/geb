{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package resources

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

/*
API:       {{API.name}}
Name:      {{RC.name}}
Route:     {{RC.route}}
Resource:  {{RC.resource}}
Path:      {{RC.path}}
Parent:    {{RC.parent}}
*/


{{#each methods}}

{{#with . as |M|}}
// {{upper M.method}}  {{M.input}}  ->  {{M.output}}
func Handle_{{upper M.method}}_{{RC.name}}(ctx echo.Context) error {

{{#if (ne M.input "none")}}
	{{#gettype M.input}}
		{{#if (builtin type)}}
			// Extract:
			input := ctx.QueryParam("{{name}}")
			// Validate:
			tag := "required{{#each validation}},{{.}}{{/each}}"
			err := validate.Var(input, tag)
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
	{{/gettype}}
{{/if}}

	// HOFSTADTER_START {{lower M.method}}
	// HOFSTADTER_END   {{lower M.method}}

	return c.JSON(http.StatusOK, input)

{{/with}}


{{/each}}
{{/with}}
{{/with}}

// HOFSTADTER_BELOW
