package render

import (
	"bytes"

	"github.com/pkg/errors"
)

// This function splices two hofstadter templates together.
// The "HOFSTADTER_START <some-name>" and "HOFSTADTER_END <some-name>" tags
// delimit sections of editable content in generated files.
// The "HOFSTADTER_BELOW" tag delimits the end of generated file content.
// You may add anything below this tag that you wish.
// The below tag can also be added to any generated file, irregardless
// of the template having the tag, and it will be respected by the splicingn process.
func SpliceResults(existing, rendered string) (string, error) {

	// get each by lines
	old_lines := bytes.Split([]byte(existing), []byte("\n"))
	new_lines := bytes.Split([]byte(rendered), []byte("\n"))

	// find HOFSTADTER tags and extract splices from the OLD file
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

	// Merge files while processing NEW file
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

	// If we foud the HOFSTADTER_BELOW line in the OLD file,
	// respect it in the NEW file and put in back
	if old_lpos > -1 {
		all_lines = append(all_lines, old_lines[old_lpos:]...)
	}

	// Rejoin the lines
	all_data := bytes.Join(all_lines, []byte("\n"))
	real_template := string(all_data)

	return real_template, nil
}
