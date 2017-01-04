package plan

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.com/aymerick/raymond"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/templates"
	// HOFSTADTER_END   import
)

// Name:      plan
// Namespace: engine.plan
// Version:   0.0.1

type Plan struct {
	Dsl             string                 `json:"dsl" xml:"dsl" yaml:"dsl" form:"dsl" query:"dsl" `
	Gen             string                 `json:"gen" xml:"gen" yaml:"gen" form:"gen" query:"gen" `
	File            string                 `json:"file" xml:"file" yaml:"file" form:"file" query:"file" `
	Data            map[string]interface{} `json:"data" xml:"data" yaml:"data" form:"data" query:"data" `
	Template        *raymond.Template      `json:"template" xml:"template" yaml:"template" form:"template" query:"template" `
	Outfile         string                 `json:"outfile" xml:"outfile" yaml:"outfile" form:"outfile" query:"outfile" `
	RepeatedContext interface{}            `json:"repeated-context" xml:"repeated-context" yaml:"repeated-context" form:"repeated-context" query:"repeated-context" `
}

func NewPlan() *Plan {
	return &Plan{
		Data:     map[string]interface{}{},
		Template: new(raymond.Template),
	}
	// loop over fields looking for pointers
}

// HOFSTADTER_BELOW

func MakePlans(dsl_map map[string]*dsl.Dsl, design_data map[string]interface{}) ([]Plan, error) {
	logger.Info("Planning Project")
	logger.Debug("    with...", "dsl_map", dsl_map, "design_data", design_data)

	var (
		ret, plans []Plan
		err        error
	)

	plans, err = make_types(dsl_map, design_data)
	if err != nil {
		return plans, errors.Wrap(err, "in MakePlans\n")
	}
	ret = append(ret, plans...)

	plans, err = make_packages(dsl_map, design_data)
	if err != nil {
		return plans, errors.Wrap(err, "in MakePlans\n")
	}
	ret = append(ret, plans...)

	plans, err = make_dsls(dsl_map, design_data)
	if err != nil {
		return plans, errors.Wrap(err, "in MakePlans\n")
	}
	ret = append(ret, plans...)

	return ret, nil
}

func determine_outfile_name(of_tpl_value string, tpl_data interface{}) (string, error) {
	tpl, err := raymond.Parse(of_tpl_value)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	templates.AddHelpers(tpl)
	of_name, err := tpl.Exec(tpl_data)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	return strings.ToLower(of_name), nil
}
