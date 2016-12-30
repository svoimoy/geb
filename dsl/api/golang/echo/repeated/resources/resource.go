{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package resources

import (
	"net/http"

	"github.com/labstack/echo"

	// HOFSTADTER_START import
	// HOFSTADTER_END   import
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
	// input
	{{#gettype M.input as |TYP|}}
		{{#if (builtin type)}}
			// Extract:
			input := ctx.QueryParam("{{name}}")
			// Validate:
			tag := "required{{#each validation}},{{.}}{{/each}}"
			err := validator.New().Var(input, tag)
			if err != nil {
				return err
			}

		{{else}}
			// Extract:
			// need to import the type and call pkg.New...
			// or add a field during prep to add package name
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
{{else}}
	// no input
{{/if}}


{{#if (ne M.output "none")}}
	// output
	{{#gettype M.output as |TYP|}}
		{{#if (builtin type)}}
			// builtin
			var input {{type}}
		{{else}}
			// user-defined
			{{#if (contains TYP.path ".views")}}
			{{> types/golang/view/var-new.go NAME="output" TYP=.}}
			{{else}}
			// Its a straight up type
			{{/if}}
		{{/if}}
	{{/gettype}}
{{else}}
	// no output
{{/if}}
	// HOFSTADTER_START {{lower M.method}}
	// HOFSTADTER_END   {{lower M.method}}

	return ctx.JSON(http.StatusOK, output)

}
{{/with}}

{{/each}}
{{/with}}
{{/with}}

// HOFSTADTER_BELOW
