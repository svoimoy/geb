{{#with RepeatedContext as |CTX| }}
{{#with DslContext as |API| }}
{{#if (eq CTX.parent DslContext.name)}}
package {{camel (trimto_first CTX.path "api." false)}}
{{else}}
package {{#if CTX.parent}}{{camel CTX.parent}}{{else}}unknown{{/if}}
{{/if}}

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/go-playground/validator"
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

	fmt.Println("{{M.method}} {{CTX.name}}") 

	// HOFSTADTER_START {{lower M.method}}-pre-db
	// HOFSTADTER_END   {{lower M.method}}-pre-db


	DB := databases.POSTGRES

	{{#if (eq (lower M.method) "list")}}
	DB.Find(&{{camel M.output.[0].name}})

	{{else if (eq (lower M.method) "get")}}
	DB.Where(&{{camel M.output.[0].name}}).First(&{{camel M.output.[0].name}})

	{{else if (eq (lower M.method) "put")}}
	DB.Save(&{{camel M.input.[0].name}})

	{{else if (eq (lower M.method) "post")}}
	// template needs fixing here (for resource spec)
	{{camel M.output.[0].name}}.{{camelT CTX.resource}}ID = uuid.New().String()
	DB.Create(&{{camel M.output.[0].name}})
	
	{{else if (eq (lower M.method) "delete")}}

	DB.Delete(&{{camel M.output.[0].name}})

	{{else}}
	// unknown method: {{M.method}}
	{{/if}}

	errs := DB.GetErrors()
	if len(errs) > 0 {
		logger.Error("with DB call", "method", "{{M.method}}", "resource", "{{CTX.name}}", "errors", errs)
		return errors.New("error making call to DB" + fmt.Sprint(errs))
	}


	// HOFSTADTER_START {{lower M.method}}-post-db
	// HOFSTADTER_END   {{lower M.method}}-post-db

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
