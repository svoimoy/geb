{{#with DslContext as |CTX| }}
package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

{{#each CTX.public-functions}}
func {{camelT name}}({{#if args}}{{> pkg/golang/func-args.go args ~}}{{/if}} ) {{#if return}}{{> pkg/golang/func-return.go return ~}}{{/if ~}} {
	// HOFSTADTER_START {{camelT name}}

	// HOFSTADTER_END   {{camelT name}}
}

{{/each}}

/*
{{{yaml CTX}}}
*/

{{/with}}
