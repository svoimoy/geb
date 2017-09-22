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

func NewDslConfig() *DslConfig {
	return &DslConfig{
		Paths:   []string{},
		Default: []gen.GeneratorConfig{},
	}
}

// HOFSTADTER_BELOW
