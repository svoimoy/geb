package project

import (
	// HOFSTADTER_START import
	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/plan"
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

	Config *Config `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `

	Available map[string]*dsl.Dsl `json:"available" xml:"available" yaml:"available" form:"available" query:"available" `

	Design *design.Design `json:"design" xml:"design" yaml:"design" form:"design" query:"design" `

	DslMap map[string]*dsl.Dsl `json:"dsl-map" xml:"dsl-map" yaml:"dsl-map" form:"dsl-map" query:"dsl-map" `

	Plans []plan.Plan `json:"plans" xml:"plans" yaml:"plans" form:"plans" query:"plans" `
}

// HOFSTADTER_BELOW
