package engine

import (
	"fmt"
	"io/ioutil"

	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/project"
)

func GenerateFile(designFile, templateFile, outputFile string) error {

	fmt.Printf("%s + %s => %s\n", designFile, templateFile, outputFile)

	D := design.NewDesign()
	err := D.ImportDesignFile(designFile)
	if err != nil {
		return err
	}

	T := gen.NewTemplateMap()
	err = T.ImportTemplateFile(templateFile)
	if err != nil {
		return err
	}

	result, err := project.RenderTemplate(T[templateFile], D.Custom)
	if err != nil {
		return err
	}

	if outputFile == "stdout" {
		fmt.Println(result)
		return nil
	}

	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}
