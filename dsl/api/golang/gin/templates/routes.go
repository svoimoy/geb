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
  {{#if is-uint }}
	{{name}}_int, err := strconv.ParseUint({{name}}, 10, 64)
	if err != nil {
		res := gin.H{"error": "{{name}} must be an unsigned integer"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	{{else ~}}
	// not a uint: '{{type}}'
	{{/if}}
{{/each}}
}

{{/each}}

{{/with}}

