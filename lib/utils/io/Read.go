package io

// package publicFiles

import (
	// HOFSTADTER_START import
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/clbanning/mxj"
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
func ReadAll(reader io.Reader, obj *interface{}) (contentType string, err error) {
	// HOFSTADTER_START ReadAll
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	// the following error checks are opposite the usual
	// we try from most common to least common types
	// xml needs to come first because json also seems to read it

	mv, merr := mxj.NewMapXml(data)
	if merr == nil {
		*obj = map[string]interface{}(mv)
		return "xml", nil
	}

	err = json.Unmarshal(data, obj)
	if err == nil {
		return "json", nil
	}

	err = yaml.Unmarshal(data, obj)
	if err == nil {
		return "yaml", nil
	}

	err = toml.Unmarshal(data, obj)
	if err == nil {
		return "toml", nil
	}

	return "", errors.New("unknown content type")
	// HOFSTADTER_END   ReadAll
	return
}

/*
Where's your docs doc?!
*/
func ReadFile(filepath string, obj *interface{}) (contentType string, err error) {
	// HOFSTADTER_START ReadFile
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	dot := strings.LastIndex(filepath, ".")
	ext := filepath[dot+1:]
	switch ext {

	case "json":
		err = json.Unmarshal(data, obj)
		if err != nil {
			return "", err
		}
		return "json", nil

	case "toml":
		err = toml.Unmarshal(data, obj)
		if err != nil {
			return "", err
		}
		return "toml", nil

	case "xml":
		mv, merr := mxj.NewMapXml(data)
		if merr != nil {
			return "", merr
		}
		*obj = map[string]interface{}(mv)
		return "xml", nil

	case "yaml", "yml":
		err = yaml.Unmarshal(data, obj)
		if err != nil {
			return "", err
		}
		return "yaml", nil

	default:
		return DetermineDataContentType(data)
	}

	return "", errors.New("unknown content type")
	// HOFSTADTER_END   ReadFile
	return
}

// HOFSTADTER_BELOW
