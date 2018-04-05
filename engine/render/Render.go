package render

import (
	// HOFSTADTER_START import
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/hofstadter-io/geb/engine/plan"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
const COMMENT = "//"
const HOF_TAG = "HOFST" + "ADTER_"

// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func RenderPlans(plans []plan.Plan, outputDir string) (errorReport []error) {
	// HOFSTADTER_START RenderPlans
	logger.Info("RenderPlans", "outputDir", outputDir)

	errs := []error{}
	for _, plan := range plans {
		err := RenderPlan(plan, outputDir)
		if err != nil {
			logger.Error("while rendering plan", "err", err, "plan", plan)
			err = errors.Wrapf(err, "in render.RenderPlans(), while render plan: %+v\n", plan)
			errs = append(errs, err)
		}
	}

	return errs
	// HOFSTADTER_END   RenderPlans
	return
}

/*
Where's your docs doc?!
*/
func RenderPlan(plan plan.Plan, outputDir string) (err error) {
	// HOFSTADTER_START RenderPlan

	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "in render.RenderPlans(), while geting CWD\n")
	}

	p_dir := filepath.Join(cwd, outputDir)
	logger.Debug("o_dir: " + p_dir)
	plan.Data["proj_outdir"] = p_dir

	plan.Data["dsl_reldir"] = plan.Dsl
	d_dir := filepath.Join(p_dir, plan.Dsl)
	plan.Data["dsl_fulldir"] = d_dir

	g_rel := filepath.Join(plan.Dsl, plan.Gen)
	plan.Data["gen_reldir"] = g_rel
	g_dir := filepath.Join(d_dir, plan.Gen)
	plan.Data["gen_fulldir"] = g_dir

	// get file basedir
	f_tmp := filepath.Dir(plan.Outfile)
	f_name, f_ddir := filepath.Split(f_tmp)
	plan.Data["file_name"] = f_name
	plan.Data["file_ddir"] = f_ddir
	f_rel := filepath.Join(g_rel, f_name)
	plan.Data["file_reldir"] = f_rel
	// f_dir := filepath.Dir(filepath.Join(p_dir, plan.Outfile))
	f_tdir := filepath.Join(p_dir, plan.Outfile)
	f_dir := filepath.Dir(f_tdir)
	plan.Data["file_fulldir"] = f_dir

	env_vars := make(map[string]string)
	vars := os.Environ()
	for _, v := range vars {
		flds := strings.Split(v, "=")
		key, val := flds[0], flds[1]
		env_vars[key] = val
	}

	plan.Data["ENV"] = env_vars

	if plan.TemplateContext != nil {
		plan.Data["TemplateContext"] = plan.TemplateContext
	} else {
		delete(plan.Data, "TemplateContext")
	}

	if plan.RepeatedContext != nil {
		plan.Data["RepeatedContext"] = plan.RepeatedContext
	} else {
		delete(plan.Data, "RepeatedContext")
	}

	if plan.DslContext != nil {
		plan.Data["DslContext"] = plan.DslContext
	} else {
		delete(plan.Data, "DslContext")
	}

	if plan.WhenContext != nil {
		plan.Data["WhenContext"] = plan.WhenContext
	} else {
		delete(plan.Data, "WhenContext")
	}

	// Render the template
	tpl := (*raymond.Template)(plan.Template.Template)
	result, err := tpl.Exec(plan.Data)
	if err != nil {
		return errors.Wrapf(err, "while executing template: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	result, err = formatContent(plan.Outfile, result)
	if err != nil {
		return errors.Wrapf(err, "while formatting result: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	// Write the results, splicing if needed
	// out_filename := filepath.Join(outputDir, plan.Outfile)
	err = WriteResults(plan.Outfile, outputDir, result)
	if err != nil {
		return errors.Wrapf(err, "while writing result: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	// Write the shadow, the rendered template without user modifications
	// After the Results writing
	shadow_filename := filepath.Join(".geb/shadow", plan.Outfile)
	err = WriteShadow(shadow_filename, result)
	if err != nil {
		return errors.Wrapf(err, "while writing shadow: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	logger.Info("Wrote file", "filename", plan.Outfile)

	return nil
	// HOFSTADTER_END   RenderPlan
	return
}

/*
Where's your docs doc?!
*/
func SpliceResults(existing string, rendered string) (spliced string, err error) {
	// HOFSTADTER_START SpliceResults

	// get each by lines
	old_lines := bytes.Split([]byte(existing), []byte("\n"))
	new_lines := bytes.Split([]byte(rendered), []byte("\n"))

	logger.Debug("Splice lengths", "old_lines", len(old_lines), "new_lines", len(new_lines))

	// find HOF tags and extract splices from the OLD file
	has_below := false
	old_lpos := -1
	splices := map[string][][]byte{}
	splice := ""
	for l, line := range old_lines {
		if bytes.Contains(line, []byte(HOF_TAG)) {
			// fmt.Println(string(line))
			if bytes.Contains(line, []byte(HOF_TAG+"BELOW")) {
				has_below = true
				old_lpos = l
				break
			} else if bytes.Contains(line, []byte(HOF_TAG+"START")) {
				// get last token
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice out start:", splice)
			} else if bytes.Contains(line, []byte(HOF_TAG+"END")) {
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
		if bytes.Contains(line, []byte(HOF_TAG)) {
			if bytes.Contains(line, []byte(HOF_TAG+"BELOW")) {
				break
			} else if bytes.Contains(line, []byte(HOF_TAG+"START")) {
				all_lines = append(all_lines, line)
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice in start:", splice)
			} else if bytes.Contains(line, []byte(HOF_TAG+"END")) {
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

	// If we foud the HOF..._BELOW line in the OLD file,
	// respect it in the NEW file and put in back
	if has_below && old_lpos > -1 {
		all_lines = append(all_lines, old_lines[old_lpos:]...)
	} else {
		all_lines = append(all_lines, []byte(COMMENT+" "+HOF_TAG+"BELOW\n\n"))
	}

	logger.Debug("   result length", "all_lines", len(all_lines))
	// Rejoin the lines
	all_data := bytes.Join(all_lines, []byte("\n"))
	real_template := string(all_data)

	return real_template, nil
	// HOFSTADTER_END   SpliceResults
	return
}

/*
Where's your docs doc?!
*/
func WriteResults(filename string, outdir string, content string) (err error) {
	// HOFSTADTER_START WriteResults
	out_filename := filepath.Join(outdir, filename)
	shadow_filename := filepath.Join(".geb/shadow", filename)
	tmp_filename := filepath.Join(".geb/tmp", filename)

	dir := filepath.Dir(out_filename)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}

	final_result := content
	file_existed := false

	_, serr := os.Stat(out_filename)
	if serr == nil {
		file_existed = true
		old_content, err := ioutil.ReadFile(out_filename)
		if err != nil {
			return errors.Wrap(err, "in render.WriteResults\n")
		}

		shadow_content, err := ioutil.ReadFile(shadow_filename)
		if err != nil {

			if _, ok := err.(*os.PathError); !ok {
				// there was some other error besides not finding the file
				return errors.Wrap(err, "in render.WriteResults\n")
				// otherwise, the shadow file doesn't exist
			}

		} else {
			// We have a shadow file for the template

			base := string(shadow_content)
			user := string(old_content)

			dmp := diffmatchpatch.New()
			b2c_diffs := dmp.DiffMain(base, content, true)
			b2u_diffs := dmp.DiffMain(base, user, true)

			if len(b2u_diffs) == 1 {
				// This file has not been changed by the user (or a formatter?)
				// So we can do nothing, because... already set
				// final_result = content
				// and we short circuit the other clauses

			} else if len(b2c_diffs) == 1 {
				// This file has not had it's design or templates changed since last (re)generation
				// but the user has changed it at some point along the way, in history [we may have dealt with the diff3 on this file before]
				// So lets apply the changes
				patches := dmp.PatchMake(base, b2u_diffs)
				patched_content, applied := dmp.PatchApply(patches, base)

				// need to check applied here
				for _, patch := range applied {
					if patch != true {
						return errors.Errorf("Failed to diff/patch %q\n%v\napplied: %v", filename, dmp.PatchToText(patches), applied)
					}
				}
				final_result = patched_content

			} else {
				// ugh oh, the file has been changed on both sides of the transformation...
				// the design or template since last regen, the user at some point in history

				/*
					fmt.Printf("%s\n-------------------\n%v\n%v\n%v\n%v\n\n", filename,
						len(b2c_diffs), len(b2u_diffs),
						dmp.DiffPrettyText(b2c_diffs),
						dmp.DiffPrettyText(b2u_diffs),
					)
				*/

				err = WriteShadow(tmp_filename, content)
				if err != nil {
					return errors.Wrap(err, "in render.WriteResults\n")
				}
				cmd := exec.Command("diff3", "-m", out_filename, shadow_filename, tmp_filename)
				stdoutStderr, err := cmd.CombinedOutput()
				if err != nil {
					if EE, ok := err.(*exec.ExitError); ok {
						if EE.Error() != "exit status 1" {
							fmt.Printf("Error during diff3 on %q %+v\n", filename, EE)
							// fmt.Printf("%s\n", stdoutStderr)
							return err

						} else {
							fmt.Println("MERGE CONFLICT in:", out_filename)
						}
					} else {
						return err
					}
				}

				final_result = string(stdoutStderr)

				/*
					// fall back to old method for now, but need to diff3
					spliced, err := SpliceResults(string(old_content), content)
					if err != nil {
						return errors.Wrap(err, "in render.WriteResults\n")
					}
					final_result = spliced
				*/
			}
		}

	} else {
		if _, ok := serr.(*os.PathError); !ok {
			return errors.Wrap(serr, "in render.WriteResults\n")
		}
	}

	if file_existed {
		// os.Remove(out_name)
		// fmt.Println("Maybe delete or backup file before writing")
	}
	err = ioutil.WriteFile(out_filename, []byte(final_result), 0644)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}

	return nil
	// HOFSTADTER_END   WriteResults
	return
}

/*
Where's your docs doc?!
*/
func WriteShadow(filename string, content string) (err error) {
	// HOFSTADTER_START WriteShadow
	dir := filepath.Dir(filename)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return errors.Wrap(err, "in render.WriteResults\n")
	}

	return nil
	// HOFSTADTER_END   WriteShadow
	return
}

// HOFSTADTER_BELOW

func formatContent(filename string, content string) (formatted string, err error) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".go":
		fmtd, ferr := format.Source([]byte(content))
		if ferr != nil {
			return "", errors.Wrap(ferr, "in render.WriteResults\n")
		}
		formatted = string(fmtd)
	}

	return
}
