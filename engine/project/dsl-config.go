package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

// Name:      dsl-config
// Namespace: engine.project
// Version:   0.0.1

type DslConfig struct {
	Paths   []string  `json:"paths" xml:"paths" yaml:"paths" form:"paths" query:"paths" `
	Default []GenPair `json:"default" xml:"default" yaml:"default" form:"default" query:"default" `
}

func NewDslConfig() *DslConfig {
	return &DslConfig{
		Paths:   []string{},
		Default: []GenPair{},
	}
	// loop over fields looking for pointers
}

/*
fields:
- name: paths
  required: true
  type: array:string
- name: default
  type: array:gen-pair
name: dsl-config
namespace: engine.project
version: 0.0.1

*/

// HOFSTADTER_BELOW
