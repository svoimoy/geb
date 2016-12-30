`
json:"{{name}}"
xml:"{{name}}"
yaml:"{{name}}"
form:"{{name}}"
query:"{{name}}"{{! for formatting ~}}
{{#if validation }}{{! end format fix}}
{{> validate.go}}{{else~}}{{/if}}
`
