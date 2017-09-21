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

	Designs []TemplateConfig `json:"Designs" xml:"Designs" yaml:"Designs" form:"Designs" query:"Designs" `

	Generators []GeneratorConfig `json:"Generators" xml:"Generators" yaml:"Generators" form:"Generators" query:"Generators" `
}

// HOFSTADTER_BELOW
