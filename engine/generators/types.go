package generators

import "github.com/aymerick/raymond"

type Generator struct {
	Name      string
	Type      string
	Language  string
	SrcPath   string
	DslDesign DesignData

	Imports   []string
	Templates map[string]*raymond.Template
	Partials  map[string]*raymond.Template
}
