package main

{{#with type.user}}
Name: {{name}}

{{#each public-fields}}
public: '{{.}}' {{> tags.go}}
{{/each}}

{{#each private-fields}}
private: '{{.}}'
{{/each}}

{{/with}}
