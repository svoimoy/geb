package plan

// package

import (
	// HOFSTADTER_START import
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"

	"github.com/aymerick/raymond"

	"github.com/hofstadter-io/dotpath"

	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
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
func MakePlans(dslMap map[string]*dsl.Dsl, designData map[string]interface{}) (ret []Plan, err error) {
	// HOFSTADTER_START MakePlans
	logger.Info("Planning Project")
	logger.Info("    with...", "dslMap", dslMap)

	flatland, err := flattenDesignData("", designData)
	if err != nil {
		return ret, errors.Wrap(err, "in MakePlans\n")
	}
	// logger.Debug("    and...  flatland!!", "flatland", flatland)

	for ctx_path, design := range flatland {
		ps := strings.Split(ctx_path, "/")
		f0 := ps[0]
		switch f0 {
		case "type", "pkg", "dsl":
			plans, err := makeProjectPlans(f0, design, dslMap, designData)
			if err != nil {
				return ret, errors.Wrap(err, "in MakePlans\n")
			}
			ret = append(ret, plans...)

		default:
			return nil, errors.New("unknown design type: '" + ctx_path + "' // " + f0)

		}

		logger.Debug("in MakingPlans", "ctx_path", ctx_path, "plans", len(ret))
	}

	return ret, nil
	// HOFSTADTER_END   MakePlans
	return
}

/*
Where's your docs doc?!
*/
func MakeSubdesignPlans(dslMap map[string]*dsl.Dsl, designData map[string]interface{}) (ret []Plan, err error) {
	// HOFSTADTER_START MakeSubdesignPlans
	logger.Info("Planning Subdesigns")
	logger.Info("    with...", "dslMap", dslMap)

	flatland, err := flattenDesignData("", designData)
	if err != nil {
		return ret, errors.Wrap(err, "in MakePlans\n")
	}
	// logger.Debug("    and...  flatland!!", "flatland", flatland)

	for ctx_path, design := range flatland {
		ps := strings.Split(ctx_path, "/")
		f0 := ps[0]
		switch f0 {
		case "type", "pkg", "dsl":
			plans, err := makeSubdesignPlans(f0, design, dslMap, designData)
			if err != nil {
				return ret, errors.Wrap(err, "in MakePlans\n")
			}
			ret = append(ret, plans...)

		default:
			return nil, errors.New("unknown design type: '" + ctx_path + "' // " + f0)

		}

		logger.Debug("in MakingPlans", "ctx_path", ctx_path, "plans", len(ret))
	}

	return ret, nil
	// HOFSTADTER_END   MakeSubdesignPlans
	return
}

/*
Where's your docs doc?!
*/
func makePlans(dslKey string, genKey string, ctxDir string, dslCtx interface{}, designData map[string]interface{}, D *dsl.Dsl, G *gen.Generator, R gen.TemplateConfig, makeDesign bool) (plans []Plan, err error) {
	// HOFSTADTER_START makePlans
	logger.Info("makePlans")
	logger.Debug("  context", "dsl_ctx", dslCtx)

	logger.Info("Processing Templates Field: '"+R.Name+"'", "field", R.Field, "dslCtx", dslCtx)

	// lookup the field used to fill in the template
	repeat_elems, err := dotpath.Get(R.Field, dslCtx, false)
	// meh...
	if err != nil || repeat_elems == nil {
		logger.Debug("Skipping Templates Field: '"+R.Name+"'", "err", err, "repeat_elems", repeat_elems)

		return nil, nil
	}

	logger.Debug("Doing Templates Field: '" + R.Name + "'")
	var c_slice []interface{}

	// the first two clauses ensure its an object of some sort
	// the last handles arrays, but omits the object check
	// should we be checking for objects at all?
	// or just arrays and other?
	switch M := repeat_elems.(type) {

	case map[string]interface{}:
		c_slice = append(c_slice, M)

		/*
			case map[interface{}]interface{}:
				c_slice = append(c_slice, M)
		*/

	case []interface{}:
		for _, elem := range M {
			// need to think about sub-sub-sub-[cli/api] and N-sub-[dsl]
			//
			// recursion!
			//
			// possibly flatten array levels
			if R.Flatten > 0 {

				switch M2 := elem.(type) {

				case map[string]interface{}:
					c_slice = append(c_slice, M2)

					/*
						case map[interface{}]interface{}:
							c_slice = append(c_slice, M2)
					*/

				case []interface{}:
					for _, elem2 := range M2 {
						c_slice = append(c_slice, elem2)
					}

				default:
					logger.Info("input is not a map or slice", "input", M)

				}
			} else {
				// just add the current array to c_slice
				c_slice = append(c_slice, elem)
			}
		}

	default:
		logger.Info("input is not a map or slice", "input", M)

	}

	// logger.Info("Done adding to c_slice", "c_slice", len(c_slice))

	// flattern c_slice
	// ....
	// ....
	tmp_c_slice := []interface{}{}
	for _, elem := range c_slice {
		if A, ok := elem.([]interface{}); ok {
			for _, a := range A {
				tmp_c_slice = append(tmp_c_slice, a)
			}
		} else {
			tmp_c_slice = append(tmp_c_slice, elem)
		}
	}
	c_slice = tmp_c_slice

	logger.Info("   Collection count", "collection", R.Field, "count", len(c_slice))
	// for all of the templates in the generator configuration, for this field
	for _, t_pair := range R.Templates {
		logger.Info("    Looking for repeat template: ", "t_pair", t_pair, "in", G.Templates)

		t_key := t_pair.In
		// need to look up in Templates or Designs here
		// ugly...
		T, ok := G.Templates[t_key]
		if makeDesign == true {
			logger.Info("    Looking for design template: ", "t_pair", t_pair, "in", G.Designs)
			T, ok = G.Designs[t_key]
		}
		if !ok {
			return nil, errors.New("Unknown repeat template: " + t_key)
		}
		// t_ray := (*raymond.Template)(T)
		t_ray := T
		logger.Debug("        found repeat template: ", "repeat", R.Name, "in", t_key)

		for idx, val := range c_slice {
			// needed because of range iteration behavior
			// also want to override when 'when' is found
			local_ctx := val
			var when_ctx interface{}

			// check the unless clause
			if t_pair.Unless != "" {
				logger.Info("Unless", "t_pair", t_pair)
				unless_elems, err := dotpath.Get(t_pair.Unless, val, false)
				logger.Debug("  elems", "unless_elems", unless_elems)
				if err == nil || unless_elems != nil {
					logger.Debug("Skipping TemplatePair Unless Field: '"+R.Name+"'", "unless", t_pair.Unless, "err", err, "unless_elems", unless_elems)
					continue
				}
			}

			// check the when clause
			if t_pair.When != "" {
				logger.Info("When", "t_pair", t_pair)
				when_elems, err := dotpath.Get(t_pair.When, val, false)
				logger.Debug("  elems", "when_elems", when_elems)
				if err != nil || when_elems == nil {
					logger.Debug("Skipping TemplatePair When Field: '"+R.Name+"'", "when", t_pair.When, "err", err, "when_elems", when_elems)
					continue
				}
				switch W := when_elems.(type) {
				case []interface{}:
					if len(W) == 0 {
						logger.Warn("Skipping TemplatePair When Field: (array) '"+R.Name+"'", "when", t_pair.When, "err", err, "when_elems", when_elems)
						continue
					}
				}

				if t_pair.Field != "" {
					when_elems, err = dotpath.Get(t_pair.Field, dslCtx, false)
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("err while looking up field (from design root) in template render pair:\n%#v\n", t_pair))
					}
					
				}

				logger.Debug("When is NOW")
				when_ctx = when_elems
			}

			logger.Debug("     context", "val", local_ctx, "idx", idx)

			// ************** move all this to the determine name ************//
			OF_name, err := determineOutfileName(t_pair.Out, val)
			if err != nil {
				return nil, errors.Wrap(err, "in make_dsls\n")
			}

			// check if outfile name has special key to build the path from the output root
			// basically just erase the ctxDir because the other directories are determined else where in the config
			// (template-configs, dependecies.[designs,generators])
			if strings.HasPrefix(OF_name, "OUTPUT_ROOT/") {
				ctxDir = ""
				OF_name = strings.TrimPrefix(OF_name, "OUTPUT_ROOT/")
			}

			G_key := filepath.Join(dslKey, genKey)
			if G.Config.OutputDir != "" {
				G_key = G.Config.OutputDir
			}

			outfile := filepath.Join(G_key, ctxDir, OF_name)
			logger.Info("OFNAME", "G_key", G_key, "ctx_dir", ctxDir, "OF_name", OF_name, "outfile", outfile)
			// ************** move all this to the determine name ************//

			// build up the plan data struct
			fgd := Plan{
				Dsl:      dslKey,
				Gen:      genKey,
				File:     t_key,
				Template: t_ray,
				Data:     designData,
				Outfile:  outfile,

				DslContext:      dslCtx,
				RepeatedContext: local_ctx,
				TemplateContext: local_ctx,
				WhenContext:     when_ctx,
			}
			// logger.Info("        planned repeat file: "+t_key, "index", idx)
			// logger.Debug("          data...", "fgd", fgd, "index", idx)

			// add the plan to a linear list to be rendered
			plans = append(plans, fgd)

		} // END of context loop 'c_slice'
		logger.Info("    end repeat loop: ", "repeat", R.Name, "in", t_key, "c_slice", len(c_slice))
		// logger.Debug("    end repeat loop: ", "repeat", R.Name, "in", t_key, "c_slice", c_slice)

	}
	logger.Info("return from makePlans")
	// logger.Debug("return from makePlans", "plans", plans)

	return plans, nil

	// HOFSTADTER_END   makePlans
	return
}

/*
Where's your docs doc?!
*/
func makeProjectPlans(dslType string, dslCtx interface{}, dslMap map[string]*dsl.Dsl, designData map[string]interface{}) (plans []Plan, err error) {
	// HOFSTADTER_START makeProjectPlans
	logger.Debug("makeProjectPlans start")
	// get the ctx path for later comparison against dsl
	ictx_path, err := dotpath.Get("ctx_path", dslCtx, true)
	if err != nil {
		return nil, errors.New("ctx_path not found, in make_type")
	}
	ctx_path, ok := ictx_path.(string)
	if !ok {
		return nil, errors.New("ctx_path is not a string, in make_type")
	}

	// For DSLs, we need the last field to know which dsl it is
	ctx_flds := strings.Split(ctx_path, ".")
	ctx_dir := ""
	if len(ctx_flds) > 2 {
		ctx_dir = filepath.Join(ctx_flds[1 : len(ctx_flds)-1]...)
	}
	ctx_dsl := ctx_flds[0]
	if dslType == "dsl" {
		ctx_dsl = ctx_flds[len(ctx_flds)-1]
	}

	logger.Debug("Making Dsl plan", "dslMap", dslMap, "ctx_dsl", ctx_dsl, "ctx_dir", ctx_dir)

	// Loop over DSLs in the plans
	for d_key, D := range dslMap {
		// ... comparing the dsl type to the design type
		if d_key != ctx_dsl {
			continue
		}
		logger.Info("    dsl: "+D.Config.Name, "d_key", d_key, "ctx_dsl", ctx_dsl, "ctx_path", ctx_path)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			//
			//  TEMPLATES
			//
			// Render the templates
			repeats := G.Config.TemplateConfigs
			if len(repeats) == 0 {
				logger.Debug("       skipping dsl repeat: "+D.Config.Type, "name", D.Config.Name, "repeats", repeats)
				continue
			}
			logger.Info("Templates found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing dsl repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			for _, R := range repeats {
				ps, err := makePlans(d_key, g_key, ctx_dir, dslCtx, designData, D, G, R, false)
				if err != nil {
					return nil, errors.Wrap(err, "while making project plans")
				}
				plans = append(plans, ps...)
			} // End of template processing

		} // End Generator loop

	} // End DSL loop

	// HOFSTADTER_END   makeProjectPlans
	return
}

/*
Where's your docs doc?!
*/
func makeSubdesignPlans(dslType string, dslCtx interface{}, dslMap map[string]*dsl.Dsl, designData map[string]interface{}) (plans []Plan, err error) {
	// HOFSTADTER_START makeSubdesignPlans
	logger.Debug("makeSubdesignPlans start")
	// get the ctx path for later comparison against dsl
	ictx_path, err := dotpath.Get("ctx_path", dslCtx, true)
	if err != nil {
		return nil, errors.New("ctx_path not found, in make_type")
	}
	ctx_path, ok := ictx_path.(string)
	if !ok {
		return nil, errors.New("ctx_path is not a string, in make_type")
	}

	// For DSLs, we need the last field to know which dsl it is
	ctx_flds := strings.Split(ctx_path, ".")
	ctx_dir := ""
	if len(ctx_flds) > 2 {
		ctx_dir = filepath.Join(ctx_flds[1 : len(ctx_flds)-1]...)
	}
	ctx_dsl := ctx_flds[0]
	if dslType == "dsl" {
		ctx_dsl = ctx_flds[len(ctx_flds)-1]
	}

	logger.Debug("Making Dsl plan", "dslMap", dslMap, "ctx_dsl", ctx_dsl, "ctx_dir", ctx_dir)

	// Loop over DSLs in the plans
	for d_key, D := range dslMap {
		// ... comparing the dsl type to the design type
		if d_key != ctx_dsl {
			continue
		}
		logger.Info("    dsl: "+D.Config.Name, "d_key", d_key, "ctx_dsl", ctx_dsl, "ctx_path", ctx_path)

		// Loop over each generator in the current DSL
		for g_key, G := range D.Generators {
			logger.Info("      gen: "+g_key, "gen_cfg", G.Config)

			//
			//  TEMPLATES
			//
			// Render the templates
			repeats := G.Config.Dependencies.Designs
			if len(repeats) == 0 {
				logger.Debug("       skipping dsl repeat: "+D.Config.Type, "name", D.Config.Name, "repeats", repeats)
				continue
			}
			logger.Info("Templates found in config:", "count", len(repeats), "repeats", repeats)
			logger.Info("      doing dsl repeat: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			// Render the repeated templates
			for _, R := range repeats {
				ps, err := makePlans(d_key, g_key, filepath.Join("subdesigns", ctx_dir), dslCtx, designData, D, G, R, true)
				if err != nil {
					return nil, errors.Wrap(err, "while making project plans")
				}
				plans = append(plans, ps...)
			} // End of template processing

		} // End Generator loop

	} // End DSL loop

	// HOFSTADTER_END   makeSubdesignPlans
	return
}

/*
Where's your docs doc?!
*/
func flattenDesignData(baseOutputPath string, designData interface{}) (flattened map[string]interface{}, err error) {
	// HOFSTADTER_START flattenDesignData
	flattened = map[string]interface{}{}

	switch D := designData.(type) {
	case map[string]interface{}:
		for key, val := range D {
			dsl_key := key
			if baseOutputPath != "" {
				dsl_key = strings.Join([]string{baseOutputPath, key}, "/")
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
					flattened[dsl_key] = vmap

					/*
						case map[interface{}]interface{}:
							flattened[dsl_key] = vmap
					*/

				default:
					return nil, errors.New("in flatten_design, named data is not a map")
				}

			} else {
				// otherwise recurse and combine the return into flat
				fs, err := flattenDesignData(dsl_key, val)
				if err != nil {
					return nil, errors.Wrap(err, "in flatten_design: "+key)
				}
				for k, v := range fs {
					flattened[k] = v
				}
			}
		}

	default:
		return nil, errors.New("in flatten_design, data is not a map")
	}

	return flattened, nil
	// HOFSTADTER_END   flattenDesignData
	return
}

/*
Where's your docs doc?!
*/
func determineOutfileName(outfileTemplateString string, renderingData interface{}) (outputFilename string, err error) {
	// HOFSTADTER_START determineOutfileName
	logger.Debug("outfile_name", "in", outfileTemplateString)
	rtpl, err := raymond.Parse(outfileTemplateString)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	tpl := &templates.Template{rtpl}
	templates.AddHelpersToTemplate(tpl)
	outputFilename, err = tpl.Exec(renderingData)
	if err != nil {
		return "", errors.Wrap(err, "in determine_outfile_name\n")
	}

	return outputFilename, nil
	// HOFSTADTER_END   determineOutfileName
	return
}

// HOFSTADTER_BELOW
