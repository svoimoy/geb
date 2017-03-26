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

{{#each CTX.functions as |F|}}
{{> common/golang/func/def.go FUNC=F}}
{{/each}}

{{/with}}

