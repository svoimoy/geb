package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      dependent-generators
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
type DependentGenerators struct {
	Dsl       string   `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Gen       []string `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `
	OutputDir string   `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

func NewDependentGenerators() *DependentGenerators {
	return &DependentGenerators{
		Gen: []string{},
	}
}

// HOFSTADTER_BELOW
