package project

import (
	// HOFSTADTER_START import
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"
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
	Name      string    `json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Version   string    `json:"version" xml:"version" yaml:"version" form:"version" query:"version" `
	About     string    `json:"about" xml:"about" yaml:"about" form:"about" query:"about" `
	DesignDir string    `json:"design-dir" xml:"design-dir" yaml:"design-dir" form:"design-dir" query:"design-dir" `
	OutputDir string    `json:"output-dir" xml:"output-dir" yaml:"output-dir" form:"output-dir" query:"output-dir" `
	DslConfig DslConfig `json:"dsl-config" xml:"dsl-config" yaml:"dsl-config" form:"dsl-config" query:"dsl-config" `
}

func NewConfig() *Config {
	return &Config{}
}

// HOFSTADTER_BELOW

func ReadConfigFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return ReadConfig(f)
}

func WriteConfigFile(filename string, c *Config) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return c.Write(f)
}

func ReadConfig(r io.Reader) (*Config, error) {
	c := NewConfig()
	err := c.Read(r)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) Read(r io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}

	data := buf.Bytes()
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Write(w io.Writer) error {

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.Write(data)
	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}
