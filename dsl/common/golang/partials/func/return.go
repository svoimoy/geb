({{#each . as |A|}}{{camel A.name}} {{> type/golang/type.go A}}{{#if @last}}{{else}},{{/if}}{{/each}}){{! formating... ~}}
