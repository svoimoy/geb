package utils

import (
	"github.com/pkg/errors"
	"os"
)

var known_file_types = []string{
	"geb.yml",
	"geb.yaml",
	"geb-dsl.yml",
	"geb-dsl.yaml",
	"geb-gen.yml",
	"geb-gen.yaml",
}

func LookForKnownFiles() string {
	for _, file := range known_file_types {
		// Does the file exist?
		_, err := os.Lstat(file)
		if err != nil {
			continue
		}
		return file
	}

	return ""
}
func ResolvePath(path string) (string, error) {
	// expand any environment vars in the path
	path = os.ExpandEnv(path)

	// does the file exist
	info, err := os.Lstat(path)
	if err != nil {
		return "", errors.Wrapf(err, "error lstat'n path in utils.ResolvePath\n")
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
}
