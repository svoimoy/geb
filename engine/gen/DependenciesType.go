package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Dependencies
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Dependencies struct {

	/* ORM: */

	Designs []TemplateConfig `json:"designs" xml:"designs" yaml:"designs" form:"designs" query:"designs" `

	Generators []GeneratorConfig `json:"generators" xml:"generators" yaml:"generators" form:"generators" query:"generators" `
}

// HOFSTADTER_BELOW
