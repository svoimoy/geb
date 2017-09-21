package project


import (
		// HOFSTADTER_START import
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
				DesignDir: design,
DesignPaths: []string{},
							OutputDir: .,
	}
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
