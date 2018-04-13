package engine

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/pkg/errors"

	"github.com/hofstadter-io/geb/engine/project"
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
func GenerateProject(filename string, paths []string) (err error) {
	// HOFSTADTER_START GenerateProject
	proj := project.NewProject()

	fmt.Println("Loading...")
	err = proj.Load(filename, paths)
	if err != nil {
		return errors.Wrapf(err, "While loading project: %s %v\n", filename, paths)
	}

	fmt.Println("Subdesigning...")
	errReport := proj.Subdesign()
	if len(errReport) > 0 {
		fmt.Println(errReport)
		return errors.Wrapf(nil, "While subdesigning project: %s %v\n", filename, paths)
	}

	err = proj.LoadDesignMerge(false, paths)
	// err = proj.LoadDesignMerge(true)
	if err != nil {
		return errors.Wrapf(err, "While reloading project design: %s %v\n", filename, paths)
	}
	errReport = proj.Unify()
	if len(errReport) > 0 {
		fmt.Println(errReport)
		return errors.Wrapf(nil, "While unifying project: %s %v\n", filename, paths)
	}

	fmt.Println("Planning...")
	err = proj.Plan()
	if err != nil {
		return errors.Wrapf(err, "While planing project: %s %v\n", filename, paths)
	}

	fmt.Println("Rendering...")
	err = proj.Render()
	if err != nil {
		return errors.Wrapf(err, "While rendering project: %s %v\n", filename, paths)
	}

	return nil
	// HOFSTADTER_END   GenerateProject
	return
}

// HOFSTADTER_BELOW
