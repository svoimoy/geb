package engine

// package publicFiles

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/project"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func GetProjectData(filename string, generators []string) (proj *project.Project, err error) {
	// HOFSTADTER_START GetProjectData
	proj = project.NewProject()

	fmt.Println("Loading...")
	err = proj.Load(filename, generators)
	if err != nil {
		return nil, errors.Wrapf(err, "While getting project data: %s %v\n", filename, generators)
	}

	fmt.Println("Planning...")
	err = proj.Plan()
	if err != nil {
		return nil, errors.Wrapf(err, "While getting project data: %s %v\n", filename, generators)
	}

	return proj, nil
	// HOFSTADTER_END   GetProjectData
	return
}

// HOFSTADTER_BELOW
