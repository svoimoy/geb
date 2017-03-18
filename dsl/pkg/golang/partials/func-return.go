({{#each . as |A|}}{{camel A.name}} {{A.type}}{{#if @last}}{{else}},{{/if}}{{/each}}){{! formating... ~}}
