{{#with RepeatedContext as |RC| }}
package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// Name:      {{RC.name}}
// Namespace: {{RC.namespace}}
// Version:   {{RC.version}}

type {{camelT RC.name}} struct {
{{#each RC.fields ~}}
	{{> type/golang/field.go .}}
{{/each}}
}

{{> type/golang/type/new-func.go TYP=RC}}


{{#each RC.views}}
{{#with . as |V|}}
type {{camelT RC.name}}View_{{camelT V.name}} struct {
{{#each V.fields}}{{#with . as |F|~}}
{{#if (hasprefix F.type "local")}}
{{#dotpath (trimprefix F.type "local.") RC.fields true }}
	{{> type/golang/field.go .}}
{{/dotpath}}
{{else}}
	{{> type/golang/field.go F}}
{{/if}}
{{/with}}{{/each ~}}
}

{{> type/golang/view/new-func.go TYP=RC VIEW=V}}
{{/with}}
{{/each}}

/*
{{{yaml RC}}}
*/

{{/with}}

// HOFSTADTER_BELOW

