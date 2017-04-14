package engine

// package publicFiles

import (
	// HOFSTADTER_START import
	"github.com/aymerick/raymond"

	"github.ibm.com/hofstadter-io/dotpath"
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
func GenerateAdhoc(inputData interface{}, fieldPath string, templateData string) (outputData string, err error) {
	// HOFSTADTER_START GenerateAdhoc

	data := inputData
	if fieldPath != "." {
		data, err = dotpath.Get(fieldPath, inputData, false)
		if err != nil {
			return "", err
		}
	}

	outputData, err = raymond.Render(templateData, data)
	if err != nil {
		return "", err
	}
	// HOFSTADTER_END   GenerateAdhoc
	return
}

// HOFSTADTER_BELOW
