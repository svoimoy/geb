package render
// package 

import (
	// HOFSTADTER_START import
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"

	"github.ibm.com/hofstadter-io/geb/engine/plan"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
const COMMENT = "//"

// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func RenderPlans(plans []plan.  Plan,outputDir string) (errorReport []error) {
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
func RenderPlan(plan plan.  Plan,outputDir string) (err error) {
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
	}

	if plan.RepeatedContext != nil {
		plan.Data["RepeatedContext"] = plan.RepeatedContext
	}

	if plan.DslContext != nil {
		plan.Data["DslContext"] = plan.DslContext
	}

	// Render the template
	tpl := (*raymond.Template)(plan.Template.Template)
	result, err := tpl.Exec(plan.Data)
	if err != nil {
		return errors.Wrapf(err, "while executing template: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	// Write the results, splicing if needed
	out_filename := filepath.Join(outputDir, plan.Outfile)
	err = WriteResults(out_filename, result)
	if err != nil {
		return errors.Wrapf(err, "while executing template: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	logger.Info("Wrote file", "filename", out_filename)

	return nil
	// HOFSTADTER_END   RenderPlan
	return
}
/*
Where's your docs doc?!
*/
func SpliceResults(existing string,rendered string) (spliced string,err error) {
	// HOFSTADTER_START SpliceResults

	// get each by lines
	old_lines := bytes.Split([]byte(existing), []byte("\n"))
	new_lines := bytes.Split([]byte(rendered), []byte("\n"))

	logger.Debug("Splice lengths", "old_lines", len(old_lines), "new_lines", len(new_lines))

	// find HOFSTADTER tags and extract splices from the OLD file
	has_below := false
	old_lpos := -1
	splices := map[string][][]byte{}
	splice := ""
	for l, line := range old_lines {
		if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_")) {
			// fmt.Println(string(line))
			if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_BELOW")) {
				has_below = true
				old_lpos = l
				break
			} else if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_START")) {
				// get last token
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice out start:", splice)
			} else if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_END")) {
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
		if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_")) {
			if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_BELOW")) {
				break
			} else if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_START")) {
				all_lines = append(all_lines, line)
				fields := bytes.Fields(line)
				splice_name := bytes.TrimSpace(fields[len(fields)-1])
				splice = string(splice_name)
				// fmt.Println("Splice in start:", splice)
			} else if bytes.Contains(line, []byte(COMMENT+" HOFSTADTER_END")) {
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
	if has_below && old_lpos > -1 {
		all_lines = append(all_lines, old_lines[old_lpos:]...)
	} else {
		all_lines = append(all_lines, []byte(COMMENT+" HOFSTADTER_BELOW\n\n"))
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
func WriteResults(filename string,content string) (err error) {
	// HOFSTADTER_START WriteResults
	dir := filepath.Dir(filename)
	err = os.MkdirAll(dir, 0755)
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
	// HOFSTADTER_END   WriteResults
	return
}



// HOFSTADTER_BELOW
