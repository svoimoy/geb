{{#with RepeatedContext as |RC| }}
{{#with DslContext as |API| }}
package subresourcename
// package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

/*
API:       {{API.name}}
Name:      {{RC.name}}
Route:     {{RC.route}}
Resource:  {{RC.resource}}
Path:      {{RC.path}}
Parent:    {{RC.parent}}

Methods:
{{#methods}}
{{#if list  }}  LIST    ({{input}}) -> {{output}}{{/if}}
{{#if get   }}  GET     ({{input}}) -> {{output}}{{/if}}
{{#if put   }}  PUT     ({{input}}) -> {{output}}{{/if}}
{{#if patch }}  PATCH   ({{input}}) -> {{output}}{{/if}}
{{#if delete}}  DELETE  ({{input}}) -> {{output}}{{/if}}
{{/methods}}
/*
{{/with}}
{{/with}}

