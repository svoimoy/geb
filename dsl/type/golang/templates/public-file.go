{{#with RepeatedContext as |RC| }}
package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

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


{{#each RC.functions}}
/*
{{#if documentation}}{{ documentation }}{{else}}Where's your docs doc?!{{/if}}

In template: type/golang/template/public-file.go
also need to case receiver type name based on private or not. need to look up type though...
*/
func ({{upper (substr RC.parent 0 1)}} *{{camelT RC.parent}}) {{camelT name}}({{#if args}}{{> pkg/golang/func-args.go args ~}}{{/if}} ) {{#if return}}{{> pkg/golang/func-return.go return ~}}{{/if ~}} {
	// HOFSTADTER_START {{camelT name}}

	// HOFSTADTER_END   {{camelT name}}
	return
}

{{/each}}
{{/with}}

// HOFSTADTER_BELOW

