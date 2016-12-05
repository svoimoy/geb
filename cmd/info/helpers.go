package cmd_info

import (
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

func look_for_file() string {
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
