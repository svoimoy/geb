{{#with RepeatedContext as |RC| }}
{{#if namespace}}
package {{#each (split RC.namespace ".")}}{{#if @last }}{{.}}{{/if}}{{/each}}
{{else}}
package types
{{/if}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:      {{RC.name}}
// Namespace: {{RC.namespace}}
// Version:   {{RC.version}}

type {{camelT RC.name}} struct {

{{#each RC.fields}}
	{{>field.go .}}
{{/each}}

}


/*

{{{yaml .}}}

*/

{{/with}}

// HOFSTADTER_BELOW

