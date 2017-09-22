package dsl

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io/ioutil"

	"github.com/ghodss/yaml"
	// HOFSTADTER_END   import
)

/*
Name:      Config
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

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
