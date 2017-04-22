package project

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/gen"
	// HOFSTADTER_END   import
)

/*
Name:      dsl-config
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
type DslConfig struct {
	Paths   []string              `json:"paths" xml:"paths" yaml:"paths" form:"paths" query:"paths" `
	Default []gen.GeneratorConfig `json:"default" xml:"default" yaml:"default" form:"default" query:"default" `
}

func NewDslConfig() *DslConfig {
	return &DslConfig{
		Paths:   []string{},
		Default: []gen.GeneratorConfig{},
	}
}

// HOFSTADTER_BELOW
