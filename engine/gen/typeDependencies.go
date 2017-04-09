package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      dependencies
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
type Dependencies struct {
	Designs    []DependentDesigns `json:"designs" xml:"designs" yaml:"designs" form:"designs" query:"designs" `
	Generators []GeneratorConfig  `json:"generators" xml:"generators" yaml:"generators" form:"generators" query:"generators" `
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		Designs:    []DependentDesigns{},
		Generators: []GeneratorConfig{},
	}
}

// HOFSTADTER_BELOW
