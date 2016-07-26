// This package implements a CLI hofstadter DSL
package design

import (
	. "github.com/hofstadter-io/geb-dsl-cli/dsl"
)

var _ = CLI("geb", func() {
	ShortHelp("The hofstadter meta-tool called geb.")

})
