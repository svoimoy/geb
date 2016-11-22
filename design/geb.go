package design

import (
	. "github.com/hofstadter-io/geb-dsl-cli/dsl"
)

var _ = CLI("geb", func() {
	ShortHelp("The hofstadter meta-tool, geb. (After Gödoel, Escher, Bach;).")

	Command("gen", func() {
		LongName("generate")
		ShortHelp("Generate from a design.")
		Flag("design", func() {
			ShortName("D")
			ShortHelp("Path to a design folder.")
			Default("./design")
		})
		Flag("output", func() {
			ShortName("O")
			ShortHelp("Path to the output folder.")
			Default(".")
		})

		Arg("generators", String, func() {
			ShortHelp("A list of generators to process the design with.")
		})
	})

	Command("qgen", func() {
		LongName("quick-generate")
		ShortHelp("Generate from an input and template to make an output.")
		Flag("input", func() {
			ShortName("I")
			ShortHelp("A filename or 'stdin'.")
		})
		Flag("template", func() {
			ShortName("I")
			ShortHelp("A filename for a golang template for the input to output transform.")
		})
		Flag("output", func() {
			ShortName("O")
			ShortHelp("A filename or 'stdout'.")
		})
	})

})
