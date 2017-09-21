package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      DslConfig
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type DslConfig struct {

	/* ORM: */

	Paths []string `json:"Paths" xml:"Paths" yaml:"Paths" form:"Paths" query:"Paths" `

	Default []gen.GeneratorConfig `json:"Default" xml:"Default" yaml:"Default" form:"Default" query:"Default" `
}

// HOFSTADTER_BELOW
