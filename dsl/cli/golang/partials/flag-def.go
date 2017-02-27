{{#if (eq FLAG.type "bool")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().BoolVarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}false{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "array:bool")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().BoolSliceVarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}[]bool{}{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "int")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().IntVarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "int8")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Int8VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "int32")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Int32VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "int64")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Int64VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "uint")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().UintVarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "uint8")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Uint8VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "uint16")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Uint16VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "uint32")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Uint32VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "uint64")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Uint64VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "float")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Float64VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0.0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "float32")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Float32VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0.0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "float64")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().Float64VarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", {{#if FLAG.default}}{{FLAG.default}}{{else}}0.0{{/if}}, "{{FLAG.help}}")
{{else if (eq FLAG.type "string")}}
{{CMDNAME}}Cmd.{{PERSIST}}Flags().StringVarP(&{{ FLAG.name }}{{#if (eq PERSIST "Persistent")}}P{{/if}}Flag, "{{FLAG.long}}", "{{FLAG.short}}", "{{#if FLAG.default}}{{FLAG.default}}{{/if}}", "{{FLAG.help}}")
{{else}}
/* unknown Flag type in:
{{{yaml FLAG}}}
*/
{{/if}}
viper.BindPFlag("{{FLAG.long}}", {{CMDNAME}}Cmd.{{PERSIST}}Flags().Lookup("{{FLAG.long}}"))

