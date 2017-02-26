{{#with DslContext as |API| }}
package main

import (
	"github.com/gin-gonic/gin"
)

// Name:     {{name}}
// Version:  {{version}}
// About:    {{about}}

func main() {
	router := gin.Default()

	group := router.Group("{{config.base-url}}")
	{{#each routes}}
	group.{{method}}("/{{route}}", {{route}}_{{method}}_Handler)
	{{/each}}

	r.Run() // listen and server on 0.0.0.0:8080
}
{{/with}}
