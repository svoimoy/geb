package render

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// This function writes the results of a rendered template.
// If the output file already exists, the new and old files
// are spliced together before writing out the results.
func WriteResults(filename, content string) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}

	final_result := content
	file_existed := false

	_, serr := os.Stat(filename)
	if serr == nil {
		file_existed = true
		old_content, err := ioutil.ReadFile(filename)
		if err != nil {
			return errors.Wrap(err, "in render.WriteResults\n")
		}

		spliced, err := SpliceResults(string(old_content), content)
		if err != nil {
			return errors.Wrap(err, "in render.WriteResults\n")
		}

		final_result = spliced

	} else {
		if _, ok := serr.(*os.PathError); !ok {
			return errors.Wrap(serr, "in render.WriteResults\n")
		}
	}

	if file_existed {
		// os.Remove(out_name)
		// fmt.Println("Maybe delete or backup file before writing")
	}
	err = ioutil.WriteFile(filename, []byte(final_result), 0644)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}

	return nil
}
