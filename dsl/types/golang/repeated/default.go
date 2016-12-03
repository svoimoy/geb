package main

{{#with type.count}}
// Name:    {{name}}
// Version: {{version}}

{{#each fields}}
{{#if private}}
private: '{{name}}'
{{else}}
public:  '{{name}}' {{> tags.go}}
{{/if}}
{{/each}}


{{/with}}
