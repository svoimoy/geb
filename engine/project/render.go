package project

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
	"github.com/hofstadter-io/geb/engine/gen"
)

func (P *Project) Render() error {
	logger.Warn("Rendering Project TBD")

	for _, plan := range P.Plans {
		tpl := (*raymond.Template)(plan.Template)
		result, err := tpl.Exec(plan.Data)
		if err != nil {
			return err
		}

		err = WriteResults(plan.Outfile, result)
		if err != nil {
			return err
		}

	}

	return nil
}

func RenderTemplate(template *gen.Template, design interface{}) (string, error) {
	tpl := (*raymond.Template)(template)
	return tpl.Exec(design)
}

func WriteResults(filename, content string) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	final_result := content
	file_existed := false

	_, serr := os.Stat(filename)
	if serr == nil {
		file_existed = true
		old_content, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		spliced, err := spliceResults(string(old_content), content)
		if err != nil {
			return err
		}

		final_result = spliced

	} else {
		if _, ok := serr.(*os.PathError); !ok {
			return serr
		}
	}

	if file_existed {
		// os.Remove(out_name)
		// fmt.Println("Maybe delete or backup file before writing")
	}
	err = ioutil.WriteFile(filename, []byte(final_result), 0644)
	if err != nil {
		return err
	}

	return nil
}

func spliceResults(existing, rendered string) (string, error) {

	// get each by lines
	old_lines := bytes.Split([]byte(existing), []byte("\n"))
	new_lines := bytes.Split([]byte(rendered), []byte("\n"))

	// find HOFSTADTER tags and extract splices
	old_lpos := -1
	splices := map[string][][]byte{}
	splice := ""
	for l, line := range old_lines {
		if bytes.Contains(line, []byte("HOFSTADTER_")) {
			// fmt.Println(string(line))
			if bytes.Contains(line, []byte("HOFSTADTER_BELOW")) {
				old_lpos = l
			} else if bytes.Contains(line, []byte("HOFSTADTER_START")) {
				// get last token
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice out start:", splice)
			} else if bytes.Contains(line, []byte("HOFSTADTER_END")) {
				// fmt.Println("Splice out end:", splice, len(splices[splice]))
				splice = ""
			}
		} else if splice != "" {
			// fmt.Println("Splicing Out: ", string(line))
			tmp := splices[splice]
			tmp = append(tmp, line)
			splices[splice] = tmp
		}
	}
	if splice != "" {
		return existing, errors.New("Unterminated splice: " + splice)
	}

	/*
		for key, _ := range splices {
			fmt.Println("SPLICE: ", key)
		}
	*/

	splice = ""
	all_lines := [][]byte{}
	for _, line := range new_lines {
		if bytes.Contains(line, []byte("HOFSTADTER_")) {
			if bytes.Contains(line, []byte("HOFSTADTER_BELOW")) {
				if old_lpos > -1 {
					all_lines = append(all_lines, line)
					all_lines = append(all_lines, old_lines[old_lpos+1:]...)
					break
				}
			} else if bytes.Contains(line, []byte("HOFSTADTER_START")) {
				all_lines = append(all_lines, line)
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice in start:", splice)
			} else if bytes.Contains(line, []byte("HOFSTADTER_END")) {
				all_lines = append(all_lines, splices[splice]...)
				all_lines = append(all_lines, line)
				// fmt.Println("Splice in end:", splice)
				splice = ""
			} else {
				// fmt.Println("Splicing In: ", string(line))
				all_lines = append(all_lines, line)
			}

		} else if splice == "" {
			all_lines = append(all_lines, line)
		}
	}

	all_data := bytes.Join(all_lines, []byte("\n"))
	real_template := string(all_data)

	return real_template, nil
}
