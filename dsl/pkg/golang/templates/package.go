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

{{#each CTX.public-functions as |F|}}
{{> common/golang/func/def.go FUNC=F}}
{{/each}}

{{#each CTX.private-functions as |F|}}
{{> common/golang/func/def.go PRIVATE="true" FUNC=F}}
{{/each}}

{{/with}}
