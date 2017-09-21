package dsl

import (
// HOFSTADTER_START import
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

	Config *Config `json:"Config" xml:"Config" yaml:"Config" form:"Config" query:"Config" `

	SourcePath string `json:"SourcePath" xml:"SourcePath" yaml:"SourcePath" form:"SourcePath" query:"SourcePath" `

	AvailableGenerators map[string]string `json:"AvailableGenerators" xml:"AvailableGenerators" yaml:"AvailableGenerators" form:"AvailableGenerators" query:"AvailableGenerators" `

	Generators map[string]*gen.Generator `json:"Generators" xml:"Generators" yaml:"Generators" form:"Generators" query:"Generators" `
}

// HOFSTADTER_BELOW
