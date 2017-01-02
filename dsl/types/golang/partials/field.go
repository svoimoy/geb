{{#with . as |F|}}
{{#if F.private}}
{{camel F.name}} {{>types/golang/type.go F ~}}
{{else}}
{{camelT F.name}} {{>types/golang/type.go F}} {{> types/golang/tags.go F ~}}
{{/if}}
{{/with}}
