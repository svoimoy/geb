package project

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	About   string `yaml:"about"`

	DesignDir string    `yaml:"design-dir"`
	OutputDir string    `yaml:"output-dir"`
	DslConfig DslConfig `yaml:"dsl-config"`
}

type DslConfig struct {
	Paths   []string  `yaml:"paths"`
	Default []GenPair `yaml:"default"`
}

type GenPair struct {
	Dsl string   `yaml:"dsl"`
	Gen []string `yaml:"gen"`
}

func NewConfig() *Config {
	return &Config{
		DslConfig: DslConfig{
			Paths:   []string{},
			Default: []GenPair{},
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
