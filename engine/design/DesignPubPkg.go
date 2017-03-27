package design
// package 

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func CreateFromFolder(folder string) (d *Design,err error) {
	// HOFSTADTER_START CreateFromFolder
	d = NewDesign()
	err = d.ImportDesignFolder(folder)
	if err != nil {
		return nil, errors.Wrap(err, "in design.CreateFromFolder: "+folder+"\n")
	}
	// HOFSTADTER_END   CreateFromFolder
	return
}



// HOFSTADTER_BELOW
