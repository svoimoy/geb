{{#each PARAMS as |P|}}
// extract

{{camel P.name}} := ctx.Param("{{P.name}}")

// validate that field

{{/each}}
