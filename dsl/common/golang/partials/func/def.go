/*
{{#if FUNC.documentation}}{{ FUNC.documentation }}{{else}}Where's your docs doc?!{{/if}}
*/
func {{#if RECEIVER ~}}({{substr RECEIVER 0 1}} *{{RECEIVER}}) {{/if}}{{! end of receiver ~}}
{{#if PRIVATE}}{{camel FUNC.name}}{{else}}{{camelT FUNC.name}}{{/if}}{{! end of func name ~}}
({{#if FUNC.args}}{{> common/golang/func/args.go FUNC.args ~}}{{/if}}) {{! end of func args ~}}
{{#if FUNC.return}}{{> common/golang/func/return.go FUNC.return ~}}{{/if}} {{!end of func return ~}} 
{
	// HOFSTADTER_START {{#if PRIVATE}}{{camel FUNC.name}}{{else}}{{camelT FUNC.name}}{{/if}}

	// HOFSTADTER_END   {{#if PRIVATE}}{{camel FUNC.name}}{{else}}{{camelT FUNC.name}}{{/if}}
	return
}
