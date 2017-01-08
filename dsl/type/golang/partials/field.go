{{#with . as |F|}}
{{#if F.private}}
{{camel F.name}} {{>type/golang/type.go F ~}}
{{else}}
{{camelT F.name}} {{>type/golang/type.go F}} {{> type/golang/tags.go F ~}}
{{/if}}
{{/with}}
