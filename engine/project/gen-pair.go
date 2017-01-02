package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

// Name:      gen-pair
// Namespace: engine.project
// Version:   0.0.1

type GenPair struct {
	Dsl       string   `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Gen       []string `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `
	OutputDir string   `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

func NewGenPair() *GenPair {
	return &GenPair{
		Gen: []string{},
	}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW
