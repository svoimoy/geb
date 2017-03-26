package project

import (
	// "fmt"
	"github.com/pkg/errors"

	// "github.com/kr/pretty"

	"github.ibm.com/hofstadter-io/geb/engine/unify"
)

func (P *Project) Unify() (errReport []error) {
	logger.Info("Unifying Project")
	var err error

	// Unify generators
	// logger.Info("Unifying generators")

	d := P.Design

	// Unify Design
	logger.Info("Unifying design")

	err = unify.Unify("", "proj", "", d.Proj)
	if err != nil {
		errReport = append(errReport, errors.Wrap(err, "While unifying design in : proj\n"))
	}

	err = unify.Unify("", "pkg", "", d.Pkg)
	if err != nil {
		errReport = append(errReport, errors.Wrap(err, "While unifying design in : pkg\n"))
	}

	err = unify.Unify("", "type", "", d.Type)
	if err != nil {
		errReport = append(errReport, errors.Wrap(err, "While unifying design in : type\n"))
	}

	err = unify.Unify("", "dsl", "", d.Dsl)
	if err != nil {
		errReport = append(errReport, errors.Wrap(err, "While unifying design in : dsl\n"))
	}

	err = unify.Unify("", "custom", "", d.Custom)
	if err != nil {
		errReport = append(errReport, errors.Wrap(err, "While unifying design in : custom\n"))
	}
	P.Design = d

	// dstr := fmt.Sprintf("%# v\n\n", pretty.Formatter(P.Design))
	// fmt.Println(dstr)
	logger.Debug("Project Unified", "design", P.Design)

	return 

}
