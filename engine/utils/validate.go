package utils

import (
	"os"
)

func AccumErrs(existing, errs []error) []error {
	if len(errs) == 0 {
		return existing
	}
	if len(existing) == 0 {
		return errs
	}
	return append(existing, errs...)
}

func FileExists(filename string) error {
	_, err := os.Lstat(filename)
	return err
}
