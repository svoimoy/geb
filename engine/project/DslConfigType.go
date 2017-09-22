package project

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/gen"
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

	Paths []string `json:"paths" xml:"paths" yaml:"paths" form:"paths" query:"paths" `

	Default []gen.GeneratorConfig `json:"default" xml:"default" yaml:"default" form:"default" query:"default" `
}

// HOFSTADTER_BELOW
