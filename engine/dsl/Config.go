package dsl

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io/ioutil"

	"github.com/ghodss/yaml"
	// HOFSTADTER_END   import
)

/*
Name:      config
About:
*/

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
type Config struct {
	Name    string                 `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Version string                 `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `
	About   string                 `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `
	Type    string                 `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
	Spec    map[string]interface{} `json:"spec" xml:"spec" yaml:"spec" form:"spec" query:"spec" `
}

func NewConfig() *Config {
	return &Config{
		Spec: map[string]interface{}{},
	}
}

// HOFSTADTER_BELOW

func ReadConfigFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading dsl config file: (readfile) %s\n", filename)
	}

	c := NewConfig()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading dsl config file: (unmarshal) %s\n", filename)
	}

	return c, nil
}
