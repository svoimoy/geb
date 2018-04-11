package engine

import (
	// HOFSTADTER_START import
	"fmt"
	"io/ioutil"

	"github.com/aymerick/raymond"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/templates"
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
func GenerateFolder(designFolder string, templateFolder string, outputFolder string) (err error) {
	// HOFSTADTER_START GenerateFolder

	// HOFSTADTER_END   GenerateFolder
	return
}

/*
Where's your docs doc?!
*/
func GenerateFile(designFile string, templateFile string, outputFile string) (err error) {
	// HOFSTADTER_START GenerateFile
	fmt.Printf("%s + %s => %s\n", designFile, templateFile, outputFile)

	D := design.NewDesign()
	err = D.ImportDesignFile(designFile)
	if err != nil {
		return errors.Wrapf(err, "While importing design file: %s %s %s\n", designFile, templateFile, outputFile)
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
	// HOFSTADTER_END   GenerateFile
	return
}

func GenerateFileWithData(data map[string]interface{}, templateData string, outputFile string) (err error) {
	// HOFSTADTER_START GenerateFile
	// fmt.Printf("%+v\n + %s => %s\n", data, templateFile, outputFile)

	// parse template
	tpl, err := raymond.Parse(templateData)
	if err != nil {
		return errors.Wrapf(err, "While generating file: %s\n", outputFile)
	}

	templates.AddHelpersToRaymond(&templates.Template{Template: tpl})

	result, err := tpl.Exec(data)
	if err != nil {
		return errors.Wrapf(err, "While generating file: %s\n", outputFile)
	}

	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return errors.Wrapf(err, "While generating file:  %s\n", outputFile)
	}

	return nil
	// HOFSTADTER_END   GenerateFile
	return
}

// HOFSTADTER_BELOW
