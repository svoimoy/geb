package project

import (
	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
)

type Project struct {
	// Read from project directories
	Config *Config

	// Design Data + Dsl/Generators
	Design *design.Design
	DslMap map[string]*dsl.Dsl

	// Rendering Plans
	Plans []FileGenData
}

func NewProject() *Project {
	return &Project{
		DslMap: map[string]*dsl.Dsl{},
	}
}
