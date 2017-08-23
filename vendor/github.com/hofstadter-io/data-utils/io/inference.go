package io

import (
	// HOFSTADTER_START import
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/naoina/toml"
	"github.com/hofstadter-io/hof-lang/lib/parser"
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
func InferDataContentType(data []byte) (contentType string, err error) {
	// HOFSTADTER_START InferDataContentType

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

	err = xml.Unmarshal(data, &obj)
	if err == nil {
		return "yaml", nil
	}

	err = toml.Unmarshal(data, &obj)
	if err == nil {
		return "toml", nil
	}

	_, err = parser.ParseReader("", bytes.NewReader(data))
	if err == nil {
		return "hof", nil
	}

	return "", errors.New("[IDCT] unknown content type")

	// HOFSTADTER_END   InferDataContentType
	return
}

/*
Where's your docs doc?!
*/
func InferFileContentType(filename string) (contentType string, err error) {
	// HOFSTADTER_START InferFileContentType

	// assume files have correct extensions
	// TODO use 'filepath.Ext()'
	ext := filepath.Ext(filename)[1:]
	switch ext {

	case "json":
		return "json", nil

	case "toml":
		return "toml", nil

	case "yaml", "yml":
		return "yaml", nil

	case "xml":
		return "xml", nil

	case "hof":
		return "hof", nil

	default:
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		return InferDataContentType(data)
	}

	// HOFSTADTER_END   InferFileContentType
	return
}

// HOFSTADTER_BELOW
