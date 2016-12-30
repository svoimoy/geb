{{#with . as |F|}}
{{#if F.private}}
{{camel F.name}} {{>types/golang/type.go F.type ~}}
{{else}}
{{camelT F.name}} {{>types/golang/type.go F.type}} {{> types/golang/tags.go F ~}}
{{/if}}
{{/with}}
