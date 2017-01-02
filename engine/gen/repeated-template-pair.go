package gen

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

// Name:      repeated-template-pair
// Namespace: engine.gen
// Version:   0.0.1

type RepeatedTemplatePair struct {
	In  string `json:"in" xml:"in" yaml:"in" form:"in" query:"in" `
	Out string `json:"out" xml:"out" yaml:"out" form:"out" query:"out" `
}

func NewRepeatedTemplatePair() *RepeatedTemplatePair {
	return &RepeatedTemplatePair{}
	// loop over fields looking for pointers
}

/*
fields:
- name: in
  required: true
  type: string
- name: out
  required: true
  type: string
name: repeated-template-pair
namespace: engine.gen
version: 0.0.1

*/

// HOFSTADTER_BELOW
