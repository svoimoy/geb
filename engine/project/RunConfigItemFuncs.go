package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      RunConfigItem
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

func NewRunConfigItem() *RunConfigItem {
	return &RunConfigItem{
		Args: []string{},
		Env:  []string{},
	}
}

// HOFSTADTER_BELOW
