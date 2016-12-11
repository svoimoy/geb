package utils

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
)

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return errors.Wrap(err, "while copying file")
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return errors.Wrap(err, "while copying file")
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return errors.Wrap(err, "while copying dir")
	}

	if !fi.IsDir() {
		return errors.New("Source is not a directory")
	}

	// ensure dest dir does not already exist
	/*
		_, err = os.Open(dest)
		if !os.IsNotExist(err) {
			return errors.New("Destination already exists")
		}
	*/

	// create dest dir

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return errors.Wrap(err, "while copying dir")
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
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
	return
}
