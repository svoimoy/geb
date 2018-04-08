package add

import (
// HOFSTADTER_START import
// HOFSTADTER_END   import
)

/*
Name:      Options
About:
*/

// HOFSTADTER_START start
// HOFSTADTER_END   start

/*
Where's your docs doc?!
*/
type Options struct {

	/* ORM: */

	Url string `json:"Url" xml:"Url" yaml:"Url" form:"Url" query:"Url" `

	Branch string `json:"Branch" xml:"Branch" yaml:"Branch" form:"Branch" query:"Branch" `

	Tag string `json:"Tag" xml:"Tag" yaml:"Tag" form:"Tag" query:"Tag" `

	Commit string `json:"Commit" xml:"Commit" yaml:"Commit" form:"Commit" query:"Commit" `

	Location string `json:"Location" xml:"Location" yaml:"Location" form:"Location" query:"Location" `

	Global bool `json:"Global" xml:"Global" yaml:"Global" form:"Global" query:"Global" `
}

// HOFSTADTER_BELOW
