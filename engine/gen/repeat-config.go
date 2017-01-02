package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

// Name:      repeat-config
// Namespace: engine.gen
// Version:   0.0.1

type RepeatConfig struct {
	Name      string                 ` json:"name" xml:"name" yaml:"name" form:"name" query:"name" `
	Field     string                 ` json:"field" xml:"field" yaml:"field" form:"field" query:"field" `
	Templates []RepeatedTemplatePair ` json:"templates" xml:"templates" yaml:"templates" form:"templates" query:"templates" `
}

func NewRepeatConfig() *RepeatConfig {
	return &RepeatConfig{}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW
