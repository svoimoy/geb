{{#with dsl.api}}
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

{{#each routes}}
func {{route}}_{{method}}_Handler(ctx *gin.Context) {
{{#each params}}
	{{name}} := c.Query("{{name}}")
	{{#if required }}
	if {{name}} == "" {
		res := gin.H{"error": "missing {{name}} in request"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	{{/if}}

	{{#if type }} {{> (concat3 "parse/" type ".go") }} {{/if}}

{{/each}}
}

{{/each}}

{{/with}}

