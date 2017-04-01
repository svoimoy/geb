{{#with RepeatedContext as |CTX| }}
{{#with DslContext as |API| }}
{{#if (eq CTX.parent DslContext.name)}}
package {{camel (trimto_first CTX.path "api." false)}}
{{else}}
package {{#if CTX.parent}}{{camel CTX.parent}}{{else}}unknown{{/if}}
{{/if}}

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/go-playground/validator"

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

var {{camel CTX.name}}Validate = validator.New()


{{#each methods}}

{{#with . as |M|}}
/*
{{#if M.documentation}}{{ M.documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
func Handle_{{upper M.method}}_{{camelT CTX.name}}(ctx echo.Context) error {
	var err error

	// input
	{{#if M.path-params}}
	// path param
		{{> api/golang/echo/input/path-params.go PARAMS=M.path-params }}
	{{/if}}

	{{#if M.input}}
	// START binding input to query/form/body params
	// Initialize
	{{#with M.input.[0] as |IN|}}
		{{> type/golang/var-new-type.go TYP=IN }}

		err = ctx.Bind(&{{camel IN.name}})
		if err != nil {
			return err
		}
		err = {{camel CTX.name}}Validate.Struct({{camel IN.name}})
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				return ctx.JSON(http.StatusBadRequest, map[string]interface{} {
					"errors": err,
					"type": "invalid",
				})
			}
			if _, ok := err.(*validator.ValidationErrors); ok {
				return ctx.JSON(http.StatusBadRequest, map[string]interface{} {
					"errors": err,
					"type": "validation",
				})
			}
			return err
		}
	{{/with}}
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
		err = {{camel CTX.name}}Validate.Var({{camel P.name}}, {{camel P.name}}_tag)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{} {
				"errors": err,
				"type": "internal",
			})
		}
	{{else}}
	   // Only built in types are supported in query/form params. Use 'input' option on the resource.method
	{{/if}}

	{{/params}}
	// END query/form/body params
	{{/if}}


{{#if M.output}}
{{#with M.output.[0] as |OUT|}}
	// OUTPUT
	{{> type/golang/var-new-type.go TYP=OUT }}
{{/with}}
{{else}}
	// NO OUTPUT
	var output struct{}{}
{{/if}}

	// HOFSTADTER_START {{lower M.method}}
	// HOFSTADTER_END   {{lower M.method}}

	// return the output response
{{#if M.output}}
{{#with M.output.[0] as |OUT|}}
	return ctx.JSON(http.StatusOK, {{camel OUT.name}})
{{/with}}
{{else}}
	return ctx.JSON(http.StatusOK, output)
{{/if}}
	return err // hacky...
}
{{/with}}

{{/each}} // End resource.methods

{{/with}}
{{/with}}

// HOFSTADTER_BELOW
