package render

import (
	"github.com/pkg/errors"

	"os"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
)

// RenderPlans renders a list of plans to an output directory.
// It is a convienence wrapper for a loop around RenderPlan.
// It will continue to process plans, accumulating the errors.
func RenderPlans(plans []plan.Plan, output_dir string) []error {
	logger.Info("RenderPlans", "output_dir", output_dir)

	errs := []error{}
	for _, plan := range plans {
		err := RenderPlan(plan, output_dir)
		if err != nil {
			err = errors.Wrapf(err, "in render.RenderPlans(), while render plan: %+v\n", plan)
			errs = append(errs, err)
		}
	}

	return nil
}

func RenderPlan(plan plan.Plan, output_dir string) error {

	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "in render.RenderPlans(), while geting CWD\n")
	}

	p_dir := filepath.Join(cwd, output_dir)
	logger.Debug("o_dir: " + p_dir)
	plan.Data["proj_outdir"] = p_dir
	d_dir := filepath.Join(p_dir, plan.Dsl)
	plan.Data["dsl_basedir"] = d_dir
	g_dir := filepath.Join(d_dir, plan.Gen)
	plan.Data["gen_basedir"] = g_dir

	// get file basedir
	f_dir := filepath.Dir(plan.Outfile)
	plan.Data["file_basedir"] = f_dir

	env_vars := make(map[string]string)
	vars := os.Environ()
	for _, v := range vars {
		flds := strings.Split(v, "=")
		key, val := flds[0], flds[1]
		env_vars[key] = val
	}

	plan.Data["ENV"] = env_vars

	if plan.RepeatedContext != nil {
		plan.Data["RepeatedContext"] = plan.RepeatedContext
	}

	// Render the template
	tpl := (*raymond.Template)(plan.Template)
	result, err := tpl.Exec(plan.Data)
	if err != nil {
		return errors.Wrapf(err, "while executing template: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	// Write the results, splicing if needed
	out_filename := filepath.Join(output_dir, plan.Outfile)
	err = WriteResults(out_filename, result)
	if err != nil {
		return errors.Wrapf(err, "while executing template: %s -> %s -> %s = %s\n", plan.Dsl, plan.Gen, plan.File, plan.Outfile)
	}

	logger.Info("Wrote file", "filename", out_filename)

	return nil
}
