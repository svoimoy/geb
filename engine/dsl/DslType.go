package dsl

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/gen"
	// HOFSTADTER_END   import
)

/*
Name:      Dsl
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Dsl struct {

	/* ORM: */

	Config *Config `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `

	SourcePath string `json:"source-path" xml:"source-path" yaml:"source-path" form:"source-path" query:"source-path" `

	AvailableGenerators map[string]string `json:"available-generators" xml:"available-generators" yaml:"available-generators" form:"available-generators" query:"available-generators" `

	Generators map[string]*gen.Generator `json:"generators" xml:"generators" yaml:"generators" form:"generators" query:"generators" `
}

// HOFSTADTER_BELOW
