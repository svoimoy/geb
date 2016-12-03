{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package routes

/*
API:       {{API.name}}
Name:      {{RC.name}}
Route:     {{RC.route}}
Resource:  {{RC.resource}}
Parent:    {{RC.parent.name}}
Parent2:   {{RC.parent.parent.name}}

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

