{{#with RepeatedContext as |RC| }}
package {{camel RC.parent}}
// package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      {{RC.name}}
About:     {{RC.about}}
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const


{{#each RC.functions as |F|}}
{{#if RC.private}}
{{> common/golang/func/def.go FUNC=F RECEIVER=(camel RC.parent) }}
{{else}}
{{> common/golang/func/def.go FUNC=F RECEIVER=(camelT RC.parent) }}
{{/if}}
{{/each}}
{{/with}}

// HOFSTADTER_BELOW

