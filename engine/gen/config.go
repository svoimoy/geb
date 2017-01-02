package gen

import (
	// HOFSTADTER_START import
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v1"
	// HOFSTADTER_END   import
)

// Name:      config
// Namespace: engine.gen
// Version:   0.0.1

type Config struct {
	Name      string         `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Version   string         `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `
	About     string         `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `
	Type      string         `json:"type" xml:"type" yaml:"type" form:"type" query:"type" `
	Language  string         `json:"language" xml:"language" yaml:"language" form:"language" query:"language" `
	Repeated  []RepeatConfig `json:"repeated" xml:"repeated" yaml:"repeated" form:"repeated" query:"repeated" `
	OutputDir string         `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
}

func NewConfig() *Config {
	return &Config{
		Repeated: []RepeatConfig{},
	}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW

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
