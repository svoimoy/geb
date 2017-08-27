package utils

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"os"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
var known_file_types = []string{
	"geb.yml",
	"geb.yaml",
	"geb-dsl.yml",
	"geb-dsl.yaml",
	"geb-gen.yml",
	"geb-gen.yaml",
}

// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func LookForKnownFiles() (filename string) {
	// HOFSTADTER_START LookForKnownFiles
	for _, file := range known_file_types {
		// Does the file exist?
		_, err := os.Lstat(file)
		if err != nil {
			continue
		}
		return file
	}

	return ""
	// HOFSTADTER_END   LookForKnownFiles
	return
}

/*
Where's your docs doc?!
*/
func ResolvePath(path string) (resolvedPath string, err error) {
	// HOFSTADTER_START ResolvePath
	// expand any environment vars in the path
	path = os.ExpandEnv(path)

	// does the file exist
	info, err := os.Lstat(path)
	if err != nil {
		return "", errors.Wrapf(err, "error lstat'n path in utils.ResolvePath:"+path+"\n")
	}

	// Is it a symlink? find the real directory
	if info.Mode()&os.ModeSymlink != 0 {
		dir, err := os.Readlink(path)
		if err != nil {
			return "", errors.Wrapf(err, "in project.LoadGeneratorList\n")
		}
		path = dir
	}
	return path, nil
	// HOFSTADTER_END   ResolvePath
	return
}

/*
Where's your docs doc?!
*/
func FileExists(filename string) (err error) {
	// HOFSTADTER_START FileExists
	_, err = os.Lstat(filename)
	return err
	// HOFSTADTER_END   FileExists
	return
}

// HOFSTADTER_BELOW
func AccumErrs(existing, errs []error) []error {
	if len(errs) == 0 {
		return existing
	}
	if len(existing) == 0 {
		return errs
	}
	return append(existing, errs...)
}
