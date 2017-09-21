package design

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
func createFromFolder(folder string) (d *Design, err error) {
	// HOFSTADTER_START createFromFolder
	d = NewDesign()
	err = d.ImportDesignFolder(folder)
	if err != nil {
		return nil, errors.Wrap(err, "in design.CreateFromFolder: "+folder+"\n")
	}
	// HOFSTADTER_END   createFromFolder
	return
}

// HOFSTADTER_BELOW
