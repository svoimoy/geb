{{#with RepeatedContext as |CTX| }}
package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

func {{camel name}}({{#if CTX.args}}{{> pkg/golang/func-args.go CTX.args ~}}{{/if ~}}) {{#if CTX.return }}{{> pkg/golang/func-return.go CTX.return ~}}{{/if }} {
	// HOFSTADTER_START {{camel CTX.name}}

	// HOFSTADTER_END   {{camel CTX.name}}
}

{{/with}}

