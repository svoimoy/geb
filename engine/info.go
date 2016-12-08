package engine

import (
	"fmt"
	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/project"
)

func GetProjectData(filename string, generators []string) (*project.Project, error) {

	proj := project.NewProject()

	fmt.Println("Loading...")
	err := proj.Load(filename, generators)
	if err != nil {
		return nil, errors.Wrapf(err, "While getting project data: %s %v\n", filename, generators)
	}

	fmt.Println("Planning...")
	err = proj.Plan()
	if err != nil {
		return nil, errors.Wrapf(err, "While getting project data: %s %v\n", filename, generators)
	}

	return proj, nil
}
