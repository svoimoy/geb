package engine

import (
	// HOFSTADTER_START import
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/kr/pretty"
	"github.com/mohae/deepcopy"
	"github.com/pkg/errors"

	"github.com/hofstadter-io/dotpath"
	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/templates"
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
func GenerateMulti(inputPaths []string, templatePaths []string, repeats []string, outputPath string) (err error) {
	// HOFSTADTER_START GenerateMulti

	fmt.Println("  inputs:    ", inputPaths)
	fmt.Println("  templates: ", templatePaths)
	fmt.Println("  repeats:   ", repeats)
	fmt.Println("  output:    ", outputPath)
	fmt.Println("")

	D := design.NewDesign()
	for _, path := range inputPaths {
		fi, err := os.Stat(path)
		if err != nil {
			return errors.Wrapf(err, "While reading input: %q\n", path)
		}

		if fi.IsDir() {
			err = D.ImportExtraFolder(path)
			if err != nil {
				return errors.Wrapf(err, "While reading input folder: %q\n", path)
			}
		} else {
			err = D.ImportExtraFile(path)
			if err != nil {
				return errors.Wrapf(err, "While reading input file: %q\n", path)
			}
		}
	}

	fmt.Println("----- input -----")
	fmt.Printf("%# v\n", pretty.Formatter(D.Extra))
	fmt.Println("-----------------\n")

	T := templates.NewTemplateMap()
	for _, path := range templatePaths {
		fi, err := os.Stat(path)
		if err != nil {
			return errors.Wrapf(err, "While reading template: %q\n", path)
		}

		if fi.IsDir() {
			err = T.ImportFromFolder(path)
			if err != nil {
				return errors.Wrapf(err, "While reading template folder: %q\n", path)
			}
		} else {
			err = T.ImportTemplateFile(path)
			if err != nil {
				return errors.Wrapf(err, "While reading template file: %q\n", path)
			}
		}

	}

	fmt.Println("--- templates ---")
	for key, _ := range T {
		fmt.Println(key)
	}
	fmt.Println("-----------------\n")

	fmt.Println("---- repeats ----")
	for _, val := range repeats {
		fields := strings.Split(val, ":")
		if len(fields) != 3 {
			return errors.Errorf("Invalid repeat: %q\n", val)
		}
		dpath, infile, outfile := fields[0], fields[1], fields[2]
		fmt.Println(dpath, "+", infile, "=", outfile)

		var REPEATS []interface{}
		REPEATS = []interface{}{deepcopy.Copy(D.Extra)}
		if dpath != "." {
			repeats, err := dotpath.Get(dpath, REPEATS, false)
			if err != nil {
				fmt.Println("DPath:", dpath)
				fmt.Println("Input:", REPEATS)

				return errors.Wrapf(err, "While looking up dotpath %q in repeat %q\n", dpath, val)
			}
			switch rType := repeats.(type) {
			case []interface{}:
				REPEATS = rType

			default:
				REPEATS = []interface{}{rType}

			}
		}

		_, ok := T[infile]
		if !ok {
			return errors.Errorf("Unknown template file: %q in repeat %q\n", infile, val)
		}

		for _, repeat := range REPEATS {

			data := map[string]interface{}{
				"ROOT":   D.Extra,
				"REPEAT": repeat,
			}
			// do actual rendering here
			result, err := templates.RenderTemplate(T[infile], data)
			if err != nil {
				return errors.Wrapf(err, "While rendering repeat:  %q %q\n", infile, val)
			}

			// determine output filename
			tpl, err := raymond.Parse(outfile)
			if err != nil {
				return errors.Wrapf(err, "While determining repeat output filename:  %q %q\n", outfile, val)
			}
			Tpl := &templates.Template{tpl}
			templates.AddHelpersToRaymond(Tpl)

			outputFile, err := tpl.Exec(repeat)
			if err != nil {
				return errors.Wrapf(err, "While determining repeat output filename:  %q %q\n", outfile, val)
			}

			filename := filepath.Join(outputPath, outputFile)

			err = os.MkdirAll(filepath.Dir(filename), 0755)
			if err != nil {
				return errors.Wrapf(err, "While writing repeat file:  %q %q\n", filename, val)
			}
			err = ioutil.WriteFile(filename, []byte(result), 0644)
			if err != nil {
				return errors.Wrapf(err, "While writing repeat file:  %q %q\n", filename, val)
			}

		}

	}
	fmt.Println("-----------------\n")

	// clean up repeats after, in case of multiple uses
	for _, val := range repeats {
		filename := strings.Split(val, ":")[1]
		delete(T, filename)
	}

	fmt.Println("---- outputs ----")
	for key, _ := range T {
		fmt.Println(key)
		filename := filepath.Join(outputPath, key)

		result, err := templates.RenderTemplate(T[key], D.Extra)
		if err != nil {
			return errors.Wrapf(err, "While rendering file:  %q %q\n", key, filename)
		}

		err = os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			return errors.Wrapf(err, "While writing file:  %q %q\n", key, filename)
		}
		err = ioutil.WriteFile(filename, []byte(result), 0644)
		if err != nil {
			return errors.Wrapf(err, "While writing file:  %q %q\n", key, filename)
		}

	}
	fmt.Println("-----------------\n")

	// HOFSTADTER_END   GenerateMulti
	return
}

// HOFSTADTER_BELOW
