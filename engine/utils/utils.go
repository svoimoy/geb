package utils

import (
	"github.com/pkg/errors"
	"os"
)

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
