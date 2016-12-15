{{#with RepeatedContext as |RC| }}
package types

// Name:    {{RC.name}}
// Version: {{RC.version}}

type {{camelT RC.name}} struct {

{{#each RC.fields}}
	{{>field.go .}}
{{/each}}

}



{{/with}}

/*

{{{yaml .}}}

*/
