package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

// Name:      dsl-config
// Namespace: engine.project
// Version:   0.0.1

type DslConfig struct {
	Paths   []string  ` json:"paths" xml:"paths" yaml:"paths" form:"paths" query:"paths" `
	Default []GenPair ` json:"default" xml:"default" yaml:"default" form:"default" query:"default" `
}

func NewDslConfig() *DslConfig {
	return &DslConfig{}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW
