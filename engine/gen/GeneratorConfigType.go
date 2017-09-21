package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      GeneratorConfig
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type GeneratorConfig struct {

	/* ORM: */

	Dsl string `json:"Dsl" xml:"Dsl" yaml:"Dsl" form:"Dsl" query:"Dsl" `

	Gen []string `json:"Gen" xml:"Gen" yaml:"Gen" form:"Gen" query:"Gen" `

	OutputDir string `json:"OutputDir" xml:"OutputDir" yaml:"OutputDir" form:"OutputDir" query:"OutputDir" `
}

// HOFSTADTER_BELOW
