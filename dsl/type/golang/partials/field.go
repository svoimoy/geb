{{#if FIELD.private}}
{{camel FIELD.name}} {{>type/golang/type.go TYP=FIELD ~}}
{{else}}
{{camelT FIELD.name}} {{>type/golang/type.go TYP=FIELD}} {{> type/golang/tags.go FIELD ~}}
{{/if}}
