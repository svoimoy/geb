package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      dsl-config
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

/*
Where's your docs doc?!
*/
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

// HOFSTADTER_BELOW
