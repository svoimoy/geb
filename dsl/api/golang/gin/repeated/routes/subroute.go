{{#with RepeatedContext as |RC| }}
{{#with dsl.api as |API| }}
package routes

// API:    {{API.name}}
// Name:   {{RC.name}}
// Route:  {{RC.route}}
// Parent: {{RC.parent.name}}
// Parent: {{RC.parent.parent.name}}

{{/with}}
{{/with}}
