package utils

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
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
func CopyFile(source string, destination string) (err error) {
	// HOFSTADTER_START CopyFile
	sf, err := os.Open(source)
	if err != nil {
		return errors.Wrap(err, "while copying file")
	}
	defer sf.Close()
	df, err := os.Create(destination)
	if err != nil {
		return errors.Wrap(err, "while copying file")
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(destination, si.Mode())
		}

	}
	// HOFSTADTER_END   CopyFile
	return
}

/*
Where's your docs doc?!
*/
func CopyDir(source string, destination string) (err error) {
	// HOFSTADTER_START CopyDir
	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return errors.Wrap(err, "while copying dir")
	}

	if !fi.IsDir() {
		return errors.New("Source is not a directory")
	}

	// ensure destination dir does not already exist
	/*
		_, err = os.Open(destination)
		if !os.IsNotExist(err) {
			return errors.New("Destination already exists")
		}
	*/

	// create destination dir

	err = os.MkdirAll(destination, fi.Mode())
	if err != nil {
		return errors.Wrap(err, "while copying dir")
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := destination + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				return errors.Wrap(err, "while copying dir")
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				return errors.Wrap(err, "while copying dir")
			}
		}

	}
	// HOFSTADTER_END   CopyDir
	return
}

// HOFSTADTER_BELOW

var CopyFolder = CopyDir
