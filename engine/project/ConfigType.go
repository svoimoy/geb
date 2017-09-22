package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Config
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Config struct {

	/* ORM: */

	Name string `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `

	Version string `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `

	About string `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `

	DesignDir string `json:"design-dir" xml:"design-dir" yaml:"design-dir" form:"design-dir" query:"design-dir" `

	DesignPaths []string `json:"design-paths" xml:"design-paths" yaml:"design-paths" form:"design-paths" query:"design-paths" `

	OutputDir string `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `

	DslConfig DslConfig `json:"dsl-config" xml:"dsl-config" yaml:"dsl-config" form:"dsl-config" query:"dsl-config" `
}

// HOFSTADTER_BELOW
