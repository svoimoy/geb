package etl

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
		TemplateConfigs: []gen.TemplateConfig{},
	}
}

// HOFSTADTER_BELOW
