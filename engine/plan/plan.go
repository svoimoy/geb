package plan

import (
	// HOFSTADTER_START import
	"github.com/pkg/errors"
	"strings"

	"github.com/aymerick/raymond"

	"github.ibm.com/hofstadter-io/dotpath"
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
	DslContext      interface{}            `json:"dsl-context" xml:"dsl-context" yaml:"dsl-context" form:"dsl-context" query:"dsl-context" `
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

func flatten_design_data(base_outpath string, design_data interface{}) (map[string]interface{}, error) {
	flat := map[string]interface{}{}

	switch D := design_data.(type) {
	case map[string]interface{}:
		for key, val := range D {
			dsl_key := key
			if base_outpath != "" {
				dsl_key = strings.Join([]string{base_outpath, key}, "/")
			}

			// Try to retrieve a name from the current object
			iname, err := dotpath.Get("name", val, true)
			if err != nil {
				if !strings.Contains(err.Error(), "could not find 'name' in object") {
					return nil, errors.Wrap(err, "in flatten_design: "+key)
				}
			}

			if iname != nil {
				// If we found a name, we found a DSL, add it to flat
				_, ok := iname.(string)
				if !ok {
					return nil, errors.New("in flatten_design, dsl '" + key + "' name is not a string")
				}

				switch vmap := val.(type) {
				case map[string]interface{}:
					flat[dsl_key] = vmap

				case map[interface{}]interface{}:
					flat[dsl_key] = vmap

				default:
					return nil, errors.New("in flatten_design, named data is not a map")
				}

			} else {
				// otherwise recurse and combine the return into flat
				fs, err := flatten_design_data(dsl_key, val)
				if err != nil {
					return nil, errors.Wrap(err, "in flatten_design: "+key)
				}
				for k, v := range fs {
					flat[k] = v
				}
			}
		}

	default:
		return nil, errors.New("in flatten_design, data is not a map")
	}

	return flat, nil
}

func MakePlans(dsl_map map[string]*dsl.Dsl, design_data map[string]interface{}) ([]Plan, error) {
	logger.Info("Planning Project")
	logger.Info("    with...", "dsl_map", dsl_map)
	// fmt.Printf("%#  v", pretty.Formatter(design_data))

	var (
		ret, plans []Plan
		err        error
	)

	flatland, err := flatten_design_data("", design_data)
	if err != nil {
		return plans, errors.Wrap(err, "in MakePlans\n")
	}
	logger.Debug("    and...  flatland!!", "flatland", flatland)

	for ctx_path, design := range flatland {
		ps := strings.Split(ctx_path, "/")
		f0 := ps[0]
		switch f0 {
		case "type":
			plans, err = make_type(design, dsl_map, design_data)
			if err != nil {
				return plans, errors.Wrap(err, "in MakePlans\n")
			}
			ret = append(ret, plans...)

		case "pkg":
			plans, err = make_package(design, dsl_map, design_data)
			if err != nil {
				return plans, errors.Wrap(err, "in MakePlans\n")
			}
			ret = append(ret, plans...)

		case "dsl":
			plans, err = make_dsl(design, dsl_map, design_data)
			if err != nil {
				return plans, errors.Wrap(err, "in MakePlans\n")
			}
			ret = append(ret, plans...)

		default:
			return nil, errors.New("unknown design type: '" + ctx_path + "' // " + f0)

		}

	}

	return ret, nil
}

func determine_outfile_name(of_tpl_value string, tpl_data interface{}) (string, error) {
	logger.Debug("outfile_name", "in", of_tpl_value, "data", tpl_data)
	tpl, err := raymond.Parse(of_tpl_value)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	templates.AddHelpers(tpl)
	of_name, err := tpl.Exec(tpl_data)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	return of_name, nil
}
