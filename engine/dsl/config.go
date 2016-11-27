package dsl

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Name    string
	Version string
	About   string

	Imports []string
}

func NewConfig() *Config {
	return &Config{
		Imports: []string{},
	}
}

func ReadConfigFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := NewConfig()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
