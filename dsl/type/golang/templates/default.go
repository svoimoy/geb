{{#with RepeatedContext as |RC| }}
package {{#each (split RC.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

/*
Name:      {{RC.name}}
About:     {{RC.about}}
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
{{#if documentation}}{{ documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
{{#if RC.private}}
type {{camel RC.name}} struct {
{{else}}
type {{camelT RC.name}} struct {
{{/if}}

{{#if (eq RC.orm "sql")}}
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
{{/if}}

{{#each RC.fields as |F| ~}}
	{{> type/golang/field.go FIELD=F}}
{{/each}}

{{> type/golang/gorm-relation-fields.go TYP=RC}}

}

{{> type/golang/type/new-func.go TYP=RC}}


{{#each RC.views}}
{{#with . as |V|}}
/*
{{#if documentation}}{{ documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
{{#if V.private}}
type {{camel RC.name}}View{{camelT V.name}} struct {
{{else}}
type {{camelT RC.name}}View{{camelT V.name}} struct {
{{/if}}

{{#if (eq V.orm "sql")}}
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
{{/if}}

{{#each V.fields}}{{#with . as |F|~}}
{{#if (hasprefix F.type "local")}}
{{#dotpath (concat3 "[name==" (trimprefix F.type "local.") "]") RC.fields true }}
{{#with . as |realF|}}
	{{> type/golang/field.go FIELD=realF}}
{{/with}}
{{/dotpath}}
{{else}}
	{{> type/golang/field.go FIELD=F}}
{{/if}}
{{/with}}{{/each ~}}
}

{{> type/golang/view/new-func.go TYP=RC VIEW=V}}
{{/with}}
{{/each}}

{{#each RC.public-functions as |F|}}
{{#if RC.private}}
{{> common/golang/func/def.go FUNC=F RECEIVER=(camel RC.name) }}
{{else}}
{{> common/golang/func/def.go FUNC=F RECEIVER=(camelT RC.name) }}
{{/if}}
{{/each}}

{{#each RC.private-functions as |F|}}
{{#if RC.private}}
{{> common/golang/func/def.go PRIVATE=true FUNC=F RECEIVER=(camel RC.name) }}
{{else}}
{{> common/golang/func/def.go PRIVATE=true FUNC=F RECEIVER=(camelT RC.name) }}
{{/if}}
{{/each}}


{{/with}}

// HOFSTADTER_BELOW

