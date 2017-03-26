{{#with DslContext as |CTX| }}
package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}
// package {{camel CTX.parent}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

{{#each CTX.public-functions}}
/*
{{#if documentation}}{{ documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
func {{camelT name}}({{#if args}}{{> pkg/golang/func-args.go args ~}}{{/if}} ) {{#if return}}{{> pkg/golang/func-return.go return ~}}{{/if ~}} {
	// HOFSTADTER_START {{camelT name}}

	// HOFSTADTER_END   {{camelT name}}
	return
}

{{/each}}
{{/with}}
