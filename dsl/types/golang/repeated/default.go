{{#with RepeatedContext as |RC| }}
{{#if namespace}}
package {{#each (split RC.namespace ".")}}{{#if @last }}{{snake .}}{{/if}}{{/each}}
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

func New{{camelT RC.name}}() *{{camelT RC.name}} {
	return &{{camelT RC.name}}{}
	// loop over fields looking for pointers
}

{{#each RC.views}}
{{#with . as |V|}}
	/* View:
		{{{V}}}
	*/
type {{camelT RC.name}}View_{{camelT V.name}} struct {

{{#each V.fields}}{{#with . as |F|}}
	{{#if (hasprefix F.type "local") }}
		{{#dotpath (trimprefix F.type "local.") RC.fields }}
		{{>field.go .}}
		{{/dotpath}}
	{{else}}
		{{>field.go F}}
	{{/if}}
{{/with}}{{/each}}
	
}

func New{{camelT RC.name}}View_{{camelT V.name}}() *{{camelT RC.name}}View_{{camelT V.name}} {
	return &{{camelT RC.name}}View_{{camelT V.name}}{}
	// loop over fields looking for pointers
}
{{/with}}
{{/each}}


/*

{{{yaml .}}}

*/

{{/with}}

// HOFSTADTER_BELOW

