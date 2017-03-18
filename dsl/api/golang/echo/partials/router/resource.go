G.{{replace (upper method) "LIST" "GET" -1}}("/{{route ~}}
{{#each path-params ~}}/:{{name ~}}
{{/each}}", resources.Handle_{{upper method}}_{{camelT name}})
