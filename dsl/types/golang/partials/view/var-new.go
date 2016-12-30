var {{NAME}} {{> types/golang/modifiers.go MOD=MOD ~}}
{{#if IMPORT~}}
{{IMPORT}}.
{{else ~}}
{{> types/golang/package.go TYP ~}}
{{/if~}}
{{camelT TYP.parent}}View_{{camelT TYP.name}}
