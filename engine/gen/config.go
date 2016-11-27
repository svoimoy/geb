package gen

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type RepeatConfig struct {
	Name      string
	Field     string
	Templates []string
}

type Config struct {
	Name    string
	Version string
	About   string

	Type     string
	Language string

	Repeated []RepeatConfig
}

func NewConfig() *Config {
	return &Config{}
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
