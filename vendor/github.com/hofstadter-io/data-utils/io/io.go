package io

import (
	// HOFSTADTER_START import
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/clbanning/mxj"
	// "github.com/ghodss/yaml"
	"gopkg.in/yaml.v2"
	// yamlB "github.com/beego/goyaml2"
	"github.com/naoina/toml"
	"github.com/hofstadter-io/hof-lang/lib/ast"
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

	if bytes.Contains(data, []byte("---")) {
		ydata := bytes.Split(data, []byte("---"))

	    var yslice []interface{}
		for _, yd := range ydata {
			var yobj interface{}
			err = yaml.Unmarshal(yd, &yobj)
			if err != nil {
				return "", err
			}
			if yobj == nil {
				continue
			}
			yslice = append(yslice, yobj)
		}

		*obj = yslice
		return "yaml", nil
	} else {
		err = yaml.Unmarshal(data, obj)
		if err == nil {
			return "yaml", nil
		}
	}

	err = toml.Unmarshal(data, obj)
	if err == nil {
		return "toml", nil
	}

	result, err := parser.ParseReader("", bytes.NewReader(data))
	if err == nil {
		hofFile := result.(ast.HofFile)
		hofData, err := hofFile.ToData()
		if err != nil {
			return "", err
		}

		*obj = hofData
		return "hof", nil
	}

	return "", errors.New("unknown content type")
	// HOFSTADTER_END   ReadAll
	return
}

/*
Where's your docs doc?!
*/
func ReadFile(filename string, obj *interface{}) (contentType string, err error) {
	// HOFSTADTER_START ReadFile

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(filename)[1:]
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
		if bytes.Contains(data, []byte("---")) {
			ydata := bytes.Split(data, []byte("---"))

			var yslice []interface{}
			for _, yd := range ydata {
				var yobj interface{}
				err = yaml.Unmarshal(yd, &yobj)
				if err != nil {
					return "", err
				}
				if yobj == nil {
					continue
				}
				yslice = append(yslice, yobj)
			}

			*obj = yslice
			return "yaml", nil
		} else {

			err = yaml.Unmarshal(data, obj)

			// yobj, err := yamlB.Read(bytes.NewReader(data))
			if err == nil {
				// *obj = yobj
				return "yaml", nil
			}
		}

	case "hof":
		result, err := parser.ParseReader("", bytes.NewReader(data))
		if err != nil {
			return "", err
		}
		hofFile := result.(ast.HofFile)
		hofData, err := hofFile.ToData()
		if err != nil {
			return "", err
		}

		*obj = hofData
		return "hof", nil

	default:
		return InferDataContentType(data)
	}

	return "", errors.New("unknown content type")
	// HOFSTADTER_END   ReadFile
	return
}

// HOFSTADTER_BELOW
