package dsl

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io/ioutil"

	"gopkg.in/yaml.v1"
	// HOFSTADTER_END   import
)

// Name:      config
// Namespace: engine.dsl
// Version:   0.0.1

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
	// loop over fields looking for pointers
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
