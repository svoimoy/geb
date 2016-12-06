package engine

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/templates"
)

func GenerateFile(designFile, templateFile, outputFile string) error {

	fmt.Printf("%s + %s => %s\n", designFile, templateFile, outputFile)

	D := design.NewDesign()
	err := D.ImportDesignFile(designFile)
	if err != nil {
		return errors.Wrapf(err, "While generating file: %s %s %s\n", designFile, templateFile, outputFile)
	}

	T := templates.NewTemplateMap()
	err = T.ImportTemplateFile(templateFile)
	if err != nil {
		return errors.Wrapf(err, "While generating file: %s %s %s\n", designFile, templateFile, outputFile)
	}

	result, err := templates.RenderTemplate(T[templateFile], D.Custom)
	if err != nil {
		return errors.Wrapf(err, "While generating file: %s %s %s\n", designFile, templateFile, outputFile)
	}

	if outputFile == "stdout" {
		fmt.Println(result)
		return nil
	}

	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return errors.Wrapf(err, "While generating file:  %s %s %s\n", designFile, templateFile, outputFile)
	}

	return nil
}
