package project

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type GenConfig struct {
	Paths   []string `yaml:"paths"`
	Default []string `yaml:"default"`
}

type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	About   string `yaml:"about"`

	OutputDir  string    `yaml:"output-dir"`
	DesignDir  string    `yaml:"design-dir"`
	Generators GenConfig `yaml:"generators"`
}

func NewConfig() *Config {
	return &Config{
		Generators: GenConfig{
			Paths:   []string{},
			Default: []string{},
		},
	}
}

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
	buf.ReadFrom(r)
	data := buf.Bytes()

	err := yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Write(w io.Writer) error {

	return nil
}
