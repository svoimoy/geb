package io

// package publicFiles

import (
	// HOFSTADTER_START import
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/naoina/toml"
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
func DetermineFileContentType(filepath string) (contentType string, err error) {
	// HOFSTADTER_START DetermineFileContentType

	// assume files have correct extensions
	dot := strings.LastIndex(filepath, ".")
	ext := filepath[dot+1:]
	switch ext {

	case "json":
		return "json", nil

	case "yaml", "yml":
		return "yaml", nil

	case "toml":
		return "toml", nil

	default:
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return "", err
		}
		return DetermineDataContentType(data)
	}
	// HOFSTADTER_END   DetermineFileContentType
	return
}

/*
Where's your docs doc?!
*/
func DetermineDataContentType(data []byte) (contentType string, err error) {
	// HOFSTADTER_START DetermineDataContentType

	// TODO: look for unique symbols in the data
	// but always try to unmarshal to be sure

	var obj interface{}

	err = json.Unmarshal(data, &obj)
	if err == nil {
		return "json", nil
	}

	err = yaml.Unmarshal(data, &obj)
	if err == nil {
		return "yaml", nil
	}

	err = toml.Unmarshal(data, &obj)
	if err == nil {
		return "toml", nil
	}

	return "", errors.New("unknown content type")

	// HOFSTADTER_END   DetermineDataContentType
	return
}

// HOFSTADTER_BELOW
