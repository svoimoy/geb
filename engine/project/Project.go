package project

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"

	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
	"github.ibm.com/hofstadter-io/geb/engine/render"
	"github.ibm.com/hofstadter-io/geb/engine/unify"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import
)

/*
Name:      project
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
type Project struct {
	Config    *Config             `json:"config" xml:"config" yaml:"config" form:"config" query:"config" `
	Available map[string]*dsl.Dsl `json:"available" xml:"available" yaml:"available" form:"available" query:"available" `
	Design    *design.Design      `json:"design" xml:"design" yaml:"design" form:"design" query:"design" `
	DslMap    map[string]*dsl.Dsl `json:"dsl-map" xml:"dsl-map" yaml:"dsl-map" form:"dsl-map" query:"dsl-map" `
	Plans     []plan.Plan         `json:"plans" xml:"plans" yaml:"plans" form:"plans" query:"plans" `
}

func NewProject() *Project {
	return &Project{

		Config: NewConfig(),

		Available: map[string]*dsl.Dsl{},

		Design: design.NewDesign(),

		DslMap: map[string]*dsl.Dsl{},
		Plans:  []plan.Plan{},
	}
}

/*
Where's your docs doc?!
*/
func (P *Project) Load(filename string, generators []string) (err error) {
	// HOFSTADTER_START Load
	logger.Debug("Reading config file", "filename", filename)
	c, err := ReadConfigFile(filename)
	if err != nil {
		return errors.Wrap(err, "while reading project config file: "+filename)
	}
	P.Config = c
	logger.Info("Project Config", "config", P.Config)

	// This should probably move below the design loading
	// or perhaps we just have a step during unification to load the dependencies and such
	err = P.LoadGenerators()
	if err != nil {
		return errors.Wrap(err, "while loading generators\n")
	}

	// make sure loading designs does not depend on the generators being loaded
	d_dir := P.Config.DesignDir
	logger.Info("Reading designs", "folder", d_dir)
	d, err := design.CreateFromFolder(d_dir)
	if err != nil {
		return errors.Wrapf(err, "While reading design folder: %s\n", d_dir)
	}
	P.Design = d

	// dstr := fmt.Sprintf("%# v\n\n", pretty.Formatter(P.Design))
	// fmt.Println(dstr)
	logger.Debug("Project Design", "design", P.Design)

	// Prep the templates now that they are loaded
	P.registerPartials()
	P.addTemplateHelpers()

	return nil
	// HOFSTADTER_END   Load
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) Unify() (errorReport []error) {
	// HOFSTADTER_START Unify
	logger.Info("Unifying Project")
	var err error

	// Unify generators
	// logger.Info("Unifying generators")

	d := P.Design

	// Unify Design
	logger.Info("Unifying design")

	err = unify.Unify("", "proj", "", d.Proj)
	if err != nil {
		errorReport = append(errorReport, errors.Wrap(err, "While unifying design in : proj\n"))
	}

	err = unify.Unify("", "pkg", "", d.Pkg)
	if err != nil {
		errorReport = append(errorReport, errors.Wrap(err, "While unifying design in : pkg\n"))
	}

	err = unify.Unify("", "type", "", d.Type)
	if err != nil {
		errorReport = append(errorReport, errors.Wrap(err, "While unifying design in : type\n"))
	}

	err = unify.Unify("", "dsl", "", d.Dsl)
	if err != nil {
		errorReport = append(errorReport, errors.Wrap(err, "While unifying design in : dsl\n"))
	}

	err = unify.Unify("", "custom", "", d.Custom)
	if err != nil {
		errorReport = append(errorReport, errors.Wrap(err, "While unifying design in : custom\n"))
	}
	P.Design = d

	// dstr := fmt.Sprintf("%# v\n\n", pretty.Formatter(P.Design))
	// fmt.Println(dstr)
	logger.Debug("Project Unified", "design", P.Design)

	// HOFSTADTER_END   Unify
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) Subdesign() (errorReport []error) {
	// HOFSTADTER_START Subdesign

	//
	//
	//  this is function is basically a first pass
	//  with all of the plan functions used in ../GenProject.go
	//
	//

	//
	//
	//  this is P.Plan()
	//
	//

	// create a map for the planning process
	data := map[string]interface{}{
		"proj":   P.Design.Proj,
		"data":   P.Design.Data,
		"type":   P.Design.Type,
		"dsl":    P.Design.Dsl,
		"pkg":    P.Design.Pkg,
		"custom": P.Design.Custom,
	}

	// call the planning module (except subdesigns here)
	plans, err := plan.MakeSubdesignPlans(P.DslMap, data)
	if err != nil {
		return []error{errors.Wrap(err, "in proj.Project.Plan()\n")}
	}

	P.Plans = plans

	//
	//
	//  this is P.Render()
	//
	//

	// render the subdesigns
	errs := render.RenderPlans(P.Plans, P.Config.OutputDir)
	if len(errs) > 0 {
		fmt.Println("Errors during subdesign rendering:")
		for i, err := range errs {
			fmt.Printf("  %d) %v\n\n", i, err)
		}
		return errs
	}

	//
	//
	//  this is P.Load() of the subdesigns
	//
	//
	P.Design.ImportDesignFolder("subdesigns")

	//
	//
	//  this is P.Unify() of the design + subdesign
	//
	//
	// just call P.Unify() again
	P.Unify()

	//
	//
	//  then we are re-ready for the Plan and Render that is about to happen
	//
	//
	return nil
	// HOFSTADTER_END   Subdesign
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) Plan() (err error) {
	// HOFSTADTER_START Plan

	// create a map for the planning process
	data := map[string]interface{}{
		"proj":   P.Design.Proj,
		"data":   P.Design.Data,
		"type":   P.Design.Type,
		"dsl":    P.Design.Dsl,
		"pkg":    P.Design.Pkg,
		"custom": P.Design.Custom,
	}

	// call the planning module
	plans, err := plan.MakePlans(P.DslMap, data)
	if err != nil {
		return errors.Wrap(err, "in proj.Project.Plan()\n")
	}

	P.Plans = plans
	return nil
	// HOFSTADTER_END   Plan
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) Render() (err error) {
	// HOFSTADTER_START Render
	errs := render.RenderPlans(P.Plans, P.Config.OutputDir)
	if len(errs) > 0 {
		fmt.Println("Errors during rendering:")
		for i, err := range errs {
			fmt.Printf("  %d) %v\n\n", i, err)
		}
		return errs[0]
	}

	return nil
	// HOFSTADTER_END   Render
	return
}

/*
Where's your docs doc?!
*/
func (P *Project) FindAvailableGenerators(paths []string) (err error) {
	// HOFSTADTER_START FindAvailableGenerators
	logger.Info("Searching for Generators")

	// If no paths are provided, use those defined in the configuration
	if len(paths) == 0 {
		paths = P.Config.DslConfig.Paths
	}

	logger.Info("DSL override order (first to last):", "paths", paths)
	if P.Available == nil {
		P.Available = map[string]*dsl.Dsl{}
	}
	for _, path := range paths {

		// Resolve the path for EnvVars, symlinks, existance
		t_path, err := utils.ResolvePath(path)
		// skip it if the file does not exist
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				continue
			}
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}

			// otherwise return the error
			return errors.Wrapf(err, "in project.FindAvailGens\n")
		}
		path = t_path

		// Find out what's available
		avail, err := dsl.FindAvailable(path)
		if err != nil {
			return errors.Wrapf(err, "in proj.FindAvailGens %v\n", paths)
		}
		for key, val := range avail {
			existing, ok := P.Available[key]
			if ok {
				existing.MergeAvailable(val)
				P.Available[key] = existing
			} else {
				P.Available[key] = val
			}
		}
	}

	return nil
	// HOFSTADTER_END   FindAvailableGenerators
	return
}

// HOFSTADTER_BELOW

func New() *Project {
	return NewProject()
}
