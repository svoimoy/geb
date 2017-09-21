package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      dependencies
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewDependencies() *Dependencies {
	return &Dependencies{
		Designs:    []TemplateConfig{},
		Generators: []GeneratorConfig{},
	}
}

// HOFSTADTER_BELOW
