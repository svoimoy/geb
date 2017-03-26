{{#with RepeatedContext as |CTX| }}
package {{camel CTX.parent}}
// package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

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

/*
{{#if documentation}}{{ documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
func {{camel name}}({{#if CTX.args}}{{> pkg/golang/func-args.go CTX.args ~}}{{/if ~}}) {{#if CTX.return }}{{> pkg/golang/func-return.go CTX.return ~}}{{/if }} {
	// HOFSTADTER_START {{camel CTX.name}}

	// HOFSTADTER_END   {{camel CTX.name}}
	return
}

{{/with}}

