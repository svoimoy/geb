package project

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
	"github.ibm.com/hofstadter-io/geb/engine/render"
	// HOFSTADTER_END   import
)

// Name:      project
// Namespace: engine.project
// Version:   0.0.1

type Project struct {
	Config    *Config             ` json:"config" xml:"config" yaml:"config" form:"config" query:"config" `
	Available map[string]*dsl.Dsl ` json:"available" xml:"available" yaml:"available" form:"available" query:"available" `
	Design    *design.Design      ` json:"design" xml:"design" yaml:"design" form:"design" query:"design" `
	DslMap    map[string]*dsl.Dsl ` json:"dsl-map" xml:"dsl-map" yaml:"dsl-map" form:"dsl-map" query:"dsl-map" `
	Plans     []plan.Plan         ` json:"plans" xml:"plans" yaml:"plans" form:"plans" query:"plans" `
}

/*
func NewProject() *Project {
	return &Project{}
	// loop over fields looking for pointers
}
*/

// HOFSTADTER_BELOW

func New() *Project {
	return NewProject()
}

func NewProject() *Project {
	return &Project{
		Config:    NewConfig(),
		Available: map[string]*dsl.Dsl{},
		Design:    design.New(),
		DslMap:    map[string]*dsl.Dsl{},
	}
}

func (P *Project) Plan() error {

	// Prep for calling the planning module
	P.register_partials()
	P.add_template_helpers()

	// create a map for the planning process
	data := map[string]interface{}{
		"proj":   P.Design.Proj,
		"type":   P.Design.Type,
		"dsl":    P.Design.Dsl,
		"custom": P.Design.Custom,
	}

	// call the planning module
	plans, err := plan.MakePlans(P.DslMap, data)
	if err != nil {
		return errors.Wrap(err, "in proj.Project.Plan()\n")
	}

	P.Plans = plans
	return nil
}

func (P *Project) Render() error {
	errs := render.RenderPlans(P.Plans, P.Config.OutputDir)
	if len(errs) > 0 {
		fmt.Println("Errors during rendering:")
		for i, err := range errs {
			fmt.Printf("  %d) %v\n\n", i, err)
		}
		return errs[0]
	}

	return nil
}
