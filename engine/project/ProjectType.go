package project

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Project
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Project struct {

	/* ORM: */

	Config *Config `json:"Config" xml:"Config" yaml:"Config" form:"Config" query:"Config" `

	Available map[string]*dsl.Dsl `json:"Available" xml:"Available" yaml:"Available" form:"Available" query:"Available" `

	Design *design.Design `json:"Design" xml:"Design" yaml:"Design" form:"Design" query:"Design" `

	DslMap map[string]*dsl.Dsl `json:"DslMap" xml:"DslMap" yaml:"DslMap" form:"DslMap" query:"DslMap" `

	Plans []plan.Plan `json:"Plans" xml:"Plans" yaml:"Plans" form:"Plans" query:"Plans" `
}

// HOFSTADTER_BELOW
