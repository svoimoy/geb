{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package resouce

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
API:       {{API.name}}
Name:      {{RC.name}}
Route:     {{RC.route}}
Resource:  {{RC.resource}}
Path:      {{RC.path}}
Parent:    {{RC.parent}}
*/


{{#methods}}

{{#with list}}
// LIST  {{input}}  ->  {{output}}
func {{RC.name}}_LIST_Handler(ctx *ginContext) {
	{{#if (ne input "none")}}
	{{> input/var-def.go}}
	{{/if}}
}
{{/with}}

{{#with get   }}
// GET  {{input}}  ->  {{output}}
func {{RC.name}}_GET_Handler(ctx *ginContext) {
	{{#if (ne input "none")}}
	{{> input/var-def.go}}
	{{/if}}
}
{{/with}}

{{#with put   }}
//  PUT  {{input}}  ->  {{output}}
func {{RC.name}}_PUT_Handler(ctx *ginContext) {
	{{#if (ne input "none")}}
	{{> input/var-def.go}}
	{{/if}}
}
{{/with}}

{{#with patch }}
//  PATCH  {{input}}  ->  {{output}}
func {{RC.name}}_PATCH_Handler(ctx *ginContext) {
	{{#if (ne input "none")}}
	{{> input/var-def.go}}
	{{/if}}
}
{{/with}}

{{#with delete}}
//  DELETE  {{input}}  ->  {{output}}
func {{RC.name}}_DELETE_Handler(ctx *ginContext) {
	{{#if (ne input "none")}}
	{{> input/var-def.go}}
	{{/if}}
}
{{/with}}

{{/methods}}









{{/with}}
{{/with}}

