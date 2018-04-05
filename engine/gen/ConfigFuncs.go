package gen

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
		NewConfigs:      map[string]TemplateConfig{},
		StaticFiles:     []StaticFilesConfig{},
		TemplateConfigs: []TemplateConfig{},
	}
}

// HOFSTADTER_BELOW