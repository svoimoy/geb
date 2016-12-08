package project

import (
	"fmt"
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
	"github.ibm.com/hofstadter-io/geb/engine/render"
)

type Project struct {

	// Read from project directories
	Config    *Config
	Available map[string]*dsl.Dsl

	// Design Data + Dsl/Generators
	Design *design.Design
	DslMap map[string]*dsl.Dsl

	// Rendering Plans
	Plans []plan.Plan
}

func NewProject() *Project {
	return &Project{
		Available: map[string]*dsl.Dsl{},
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
