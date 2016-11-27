package project

import (
	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
)

type Project struct {
	// Read from project directories
	Config *Config
	Design *design.Design

	DslMap map[string]*dsl.Dsl
	Plans  []FileGenData

	// Pipelines  []Pipeline
}

func NewProject() *Project {
	return &Project{
		DslMap: map[string]*dsl.Dsl{},
	}
}
