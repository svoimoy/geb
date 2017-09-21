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

	Name string `json:"Name" xml:"Name" yaml:"Name" form:"Name" query:"Name" `

	Version string `json:"Version" xml:"Version" yaml:"Version" form:"Version" query:"Version" `

	About string `json:"About" xml:"About" yaml:"About" form:"About" query:"About" `

	DesignDir string `json:"DesignDir" xml:"DesignDir" yaml:"DesignDir" form:"DesignDir" query:"DesignDir" `

	DesignPaths []string `json:"DesignPaths" xml:"DesignPaths" yaml:"DesignPaths" form:"DesignPaths" query:"DesignPaths" `

	OutputDir string `json:"OutputDir" xml:"OutputDir" yaml:"OutputDir" form:"OutputDir" query:"OutputDir" `

	DslConfig DslConfig `json:"DslConfig" xml:"DslConfig" yaml:"DslConfig" form:"DslConfig" query:"DslConfig" `
}

// HOFSTADTER_BELOW
