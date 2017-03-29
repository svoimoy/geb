{{#with RepeatedContext as |CTX| }}
{{#with DslContext as |API| }}
package {{camel CTX.parent}}

/*
API:       {{API.name}}
Name:      {{CTX.name}}
Route:     {{CTX.route}}
Resource:  {{CTX.resource}}
Path:      {{CTX.path}}
Parent:    {{CTX.parent}}

Methods:
{{#methods}}
{{#if list  }}  LIST    ({{input}}) -> {{output}}{{/if}}
{{#if get   }}  GET     ({{input}}) -> {{output}}{{/if}}
{{#if put   }}  PUT     ({{input}}) -> {{output}}{{/if}}
{{#if patch }}  PATCH   ({{input}}) -> {{output}}{{/if}}
{{#if delete}}  DELETE  ({{input}}) -> {{output}}{{/if}}
{{/methods}}

dotpath to parent:
{{concat2 CTX.parent_path '.name'}}

pkgPath: {{CTX.pkgPath}}

- {{trimto_first CTX.pkgPath '/' false}}
- {{trimfrom_last CTX.pkgPath '/' false}}

{{trimto_first (trimfrom_last CTX.pkgPath '/' false) '/' false}}

- {{getbetween CTX.pkgPath '/' '/'}}
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
{{{yaml CTX}}}
*/

{{/with}}
{{/with}}

