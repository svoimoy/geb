{{camel R.route}}Group.{{replace (upper M.method) "LIST" "GET" -1}}("{{#each M.path-params as |P|~}}/:{{P.name ~}}
{{/each}}", {{#if (eq R.parent DslContext.name)}}resources{{else}}{{camel R.parent}}{{/if ~}}
.Handle_{{upper M.method}}_{{camelT R.name}})
