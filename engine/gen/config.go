package gen

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v1"
)

type Config struct {
	Name     string
	Version  string
	About    string
	Type     string
	Language string

	Repeated []RepeatConfig

	OutputDir string `yaml:"output-dir"`
}

type RepeatConfig struct {
	Name      string
	Field     string
	Templates []RepeatTemplatePair
}

type RepeatTemplatePair struct {
	In  string
	Out string
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfigFile(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading generator config file: (readfile) %s\n", filename)
	}

	c := NewConfig()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading generator config file: (unmarshal) %s\n", filename)
	}

	return c, nil
}
