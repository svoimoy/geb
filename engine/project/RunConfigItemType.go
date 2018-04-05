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

/*
Where's your docs doc?!
*/
type RunConfigItem struct {

	/* ORM: */

	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `

	Command string `json:"command" xml:"command" yaml:"command" form:"command" query:"command" `

	Args []string `json:"args" xml:"args" yaml:"args" form:"args" query:"args" `

	Workdir string `json:"workdir" xml:"workdir" yaml:"workdir" form:"workdir" query:"workdir" `

	Env []string `json:"env" xml:"env" yaml:"env" form:"env" query:"env" `
}

// HOFSTADTER_BELOW
