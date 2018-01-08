package main

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      context
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
type Context struct {
	Obj  interface{} `json:"obj" xml:"obj" yaml:"obj" form:"obj" query:"obj" `
	Path []string    `json:"path" xml:"path" yaml:"path" form:"path" query:"path" `
	Data interface{} `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `
}

func NewContext() *Context {
	return &Context{
		Path: []string{},
	}
}

// HOFSTADTER_BELOW
