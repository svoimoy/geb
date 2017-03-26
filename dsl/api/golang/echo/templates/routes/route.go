{{#with RepeatedContext as |RC| }}
{{#with DslContext as |API| }}
package routes
// package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
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


// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init


// Should find a way to build up errors and return all
// {{upper RC.method}}  {{RC.input}}  ->  {{RC.output}}
func Handle_{{upper RC.method}}_{{camelT RC.name}}(ctx echo.Context) error {
	// Check params
{{#params . as |P|}}

	{{#if (builtin P.type)}}
		// Extract
		{{camel P.name}} := ctx.QueryParam("{{P.name}}")
		// Validate
		{{camel P.name}}_tag := "required{{#each validation}},{{.}}{{/each}}"
		err := validator.New().Var({{camel P.name}}, {{camel P.name}}_tag)
		if err != nil {
			return err
		}
	{{else}}
			// Initialize
			{{> type/golang/var-new.go M.input }}

			// Extract:
			// need to import the type and call pkg.New...
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

	return nil
}

{{/with}}
{{/with}}
