package plan

import (
	// HOFSTADTER_START import
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"

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
		return ret, errors.Wrap(err, "in MakeSubdesignPlans\n")
	}
	// logger.Debug("    and...  flatland!!", "flatland", flatland)

	for ctx_path, design := range flatland {
		ps := strings.Split(ctx_path, "/")
		f0 := ps[0]
		switch f0 {
		case "type", "pkg", "dsl":
			plans, err := makeSubdesignPlans(f0, design, dslMap, designData)
			if err != nil {
				return ret, errors.Wrap(err, "in MakeSubdesignPlans\n")
			}
			ret = append(ret, plans...)

		default:
			return nil, errors.New("unknown design type: '" + ctx_path + "' // " + f0)

		}

		logger.Debug("in MakeSubdesignPlans", "ctx_path", ctx_path, "plans", len(ret))
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

	// lookup the field used to fill in the template, skip if err or not found
	fieldElems, err := dotpath.Get(R.Field, dslCtx, false)
	if err != nil {
		logger.Debug("Skipping Templates Config on error: '"+R.Name+"'", "err", err, "fieldElems", fieldElems)
		return nil, nil
	}
	if fieldElems == nil {
		logger.Debug("Skipping Templates Config on empty: '"+R.Name+"'", "err", err, "fieldElems", fieldElems)
		return nil, nil
	}

	logger.Debug("Doing Templates Field: '" + R.Name + "'")

	// the first two clauses ensure its an object of some sort
	// the last handles arrays, but omits the object check
	// should we be checking for objects at all?
	// or just arrays and other?
	var c_slice []interface{}
	switch M := fieldElems.(type) {

	// the fieldElems is actully just an object, so make it a single element slice
	case map[string]interface{}:
		c_slice = append(c_slice, M)

	case []interface{}:
		for _, elem := range M {

			if R.Flatten > 0 {
				flatten := R.Flatten
				// TODO test this loop, only '1' ever used anywhere
				for flatten > 0 {

					switch M2 := elem.(type) {

					case map[string]interface{}:
						c_slice = append(c_slice, M2)

					case []interface{}:
						for _, elem2 := range M2 {
							c_slice = append(c_slice, elem2)
						}

					default:
						logger.Info("input is not a map or slice", "input", M)

					}

					flatten -= 1
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

	//
	// GeneratorStatic Files
	//
	for _, cfg := range R.StaticFiles {
		for idx, val := range c_slice {
			if cfg.Unless != "" {
				found, err := testUnlessConditions(cfg.Unless, val)
				if err != nil {
					logger.Debug("skipping static files on unless error", "cfg", cfg, "err", err)
					continue
				}
				if found != nil {
					logger.Debug("skipping static files on unless hit", "cfg", cfg, "found", found)
					continue
				}
			}

			if cfg.When != "" {
				found, err := testWhenConditions(cfg.When, val)
				if err != nil {
					logger.Debug("skipping static files on when error", "cfg", cfg, "err", err)
					continue
				}
				if found == nil {
					logger.Debug("skipping static files on when miss", "cfg", cfg, "found", found)
					continue
				}
			}

			logger.Debug("     context", "val", val, "idx", idx)

			planFile := func(filename string) error {

				content, err := ioutil.ReadFile(filename)
				if err != nil {
					return err
				}

				G_key := filepath.Join(dslKey, genKey)
				if G.Config.OutputDir != "" {
					G_key = G.Config.OutputDir
				}

				outfile := filepath.Join(G_key, ctxDir, filename)
				logger.Info("OFNAME", "G_key", G_key, "filename", filename, "outfile", outfile)

				// build up the plan data struct
				fgd := Plan{
					Dsl:           dslKey,
					Gen:           genKey,
					File:          filename,
					StaticContent: content,
					Outfile:       outfile,
				}
				// logger.Info("        planned repeat file: "+t_key, "index", idx)
				// logger.Debug("          data...", "fgd", fgd, "index", idx)

				// add the plan to a linear list to be rendered
				plans = append(plans, fgd)
				return nil
			}

			// Gather the ignore file globs and sort
			ignores := []string{}
			for _, ignoreGlob := range cfg.Ignores {
				matches, err := filepath.Glob(ignoreGlob)
				if err != nil {
					logger.Error("Static Ignore Glob Error", "ignoreGlob", ignoreGlob, "dslKey", dslKey, "genKey", genKey, "error", err)
					return plans, err
				}
				ignores = append(ignores, matches...)
			}
			sort.Strings(ignores)

			// For each glob in the files list
			for _, fileGlob := range cfg.Files {

				// get the files for the glob
				matches, err := filepath.Glob(fileGlob)
				// tmp Warn logging for dev purposes
				logger.Warn("Static File Glob", "fileGlob", fileGlob, "dslKey", dslKey, "genKey", genKey)
				if err != nil {
					logger.Error("Static File Glob error", "fileGlob", fileGlob, "dslKey", dslKey, "genKey", genKey, "error", err)
					return plans, err
				}

				// for each file
				for _, match := range matches {
					// check no match to any ignoreGlobs
					idx := sort.SearchStrings(ignores, match)
					// two conditions for not matching an ignore
					if idx == len(ignores) || ignores[idx] != match {
						// if we get here, the static file nees to be generated
						perr := planFile(match)
						if perr != nil {
							logger.Error("Static File Glob error", "fileGlob", fileGlob, "dslKey", dslKey, "genKey", genKey, "error", perr)
							return plans, perr
						}
					}
					// otherwise we will ignore the file
				}

			}

		}
	}

	//
	// Generator Templates
	//

	// for all of the templates in the generator configuration, for this field
	for _, cfg := range R.Templates {
		logger.Info("    Looking for repeat template: ", "cfg", cfg, "in", G.Templates)

		t_key := cfg.In
		// need to look up in Templates or Designs here
		// ugly...
		T, ok := G.Templates[t_key]
		if makeDesign == true {
			logger.Info("    Looking for design template: ", "cfg", cfg, "in", G.Designs)
			T, ok = G.Designs[t_key]
		}
		if !ok {
			return nil, errors.New("Unknown repeat template: " + t_key)
		}
		// t_ray := (*raymond.Template)(T)
		t_ray := T
		logger.Debug("        found repeat template: ", "repeat", R.Name, "in", t_key)

		for idx, val := range c_slice {
			if cfg.Unless != "" {
				found, err := testUnlessConditions(cfg.Unless, val)
				if err != nil {
					logger.Debug("skipping rendering on unless error", "cfg", cfg, "err", err)
					continue
				}
				if found != nil {
					logger.Debug("skipping rendering on unless hit", "cfg", cfg, "found", found)
					continue
				}
			}

			var whenCtx interface{}
			if cfg.When != "" {
				found, err := testWhenConditions(cfg.When, val)
				if err != nil {
					logger.Debug("skipping rendering on when error", "cfg", cfg, "err", err)
					continue
				}
				if found == nil {
					logger.Debug("skipping rendering on when miss", "cfg", cfg, "found", found)
					continue
				}
				// use the when field by default
				whenCtx = found
			}

			// Override the when context
			if cfg.Field != "" {
				whenCtx, err = dotpath.Get(cfg.Field, dslCtx, false)
				if err != nil {
					return nil, errors.Wrap(err, fmt.Sprintf("err while looking up when override 'field' (from design root) in template render pair:\n%v\n", cfg))
				}
			}

			logger.Debug("     context", "val", val, "idx", idx)

			OF_name, err := determineOutfileName(cfg.Out, val, ctxDir)
			if err != nil {
				return nil, errors.Wrap(err, "in make_dsls\n")
			}

			G_key := filepath.Join(dslKey, genKey)
			if G.Config.OutputDir != "" {
				G_key = G.Config.OutputDir
			}

			outfile := filepath.Join(G_key, OF_name)
			logger.Info("OFNAME", "G_key", G_key, "OF_name", OF_name, "outfile", outfile)

			// needed because of range iteration behavior
			localCtx := val

			// build up the plan data struct
			fgd := Plan{
				Dsl:      dslKey,
				Gen:      genKey,
				File:     t_key,
				Template: t_ray,
				Data:     designData,
				Outfile:  outfile,

				DslContext:      dslCtx,
				RepeatedContext: localCtx,
				TemplateContext: localCtx,
				WhenContext:     whenCtx,
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
		return nil, errors.New("ctx_path not found, in makeProjectPlans")
	}
	ctx_path, ok := ictx_path.(string)
	if !ok {
		return nil, errors.New("ctx_path is not a string, in makeProjectPlans")
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
			//  STATIC FILES
			//    Render the static files
			//
			/*
				ps, err := makeGeneratorStaticFiles(d_key, g_key, D, G)
				if err != nil {
					return nil, errors.Wrap(err, "while making project plans")
				}
				plans = append(plans, ps...)
			*/

			//
			//  TEMPLATES
			//    Render the templates
			//
			tplCfgs := G.Config.TemplateConfigs
			if len(tplCfgs) == 0 {
				logger.Debug("       skipping generator templates: "+D.Config.Type, "name", D.Config.Name, "tplCfgs", tplCfgs)
				continue
			}

			logger.Info("Templates found in config:", "count", len(tplCfgs), "tplCfgs", tplCfgs)
			logger.Info("      doing generator templates: "+D.Config.Type, "name", D.Config.Name, "d_key", d_key)

			for _, R := range tplCfgs {
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
		return nil, errors.New("ctx_path not found, in makeSubdesignPlans : " + dslType)
	}
	ctx_path, ok := ictx_path.(string)
	if !ok {
		return nil, errors.New("ctx_path is not a string, in makeSubdesignPlans")
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
func determineOutfileName(outfileTemplateString string, renderingData interface{}, ctxDir string) (outputFilename string, err error) {
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

	// check if outfile name has special key to build the path from the output root
	// basically just erase the ctxDir because the other directories are determined else where in the config
	// (template-configs, dependecies.[designs,generators])

	if strings.HasPrefix(outputFilename, "OUTPUT_ROOT/") {
		outputFilename = strings.TrimPrefix(outputFilename, "OUTPUT_ROOT/")
	}

	if strings.HasPrefix(outputFilename, "SUBDESIGN_ROOT/") {
		ctxDir = "subdesigns" // is this a hard value?
		outputFilename = strings.TrimPrefix(outputFilename, "SUBDESIGN_ROOT/")
	}
	outputFilename = filepath.Join(ctxDir, outputFilename)

	return outputFilename, nil
	// HOFSTADTER_END   determineOutfileName
	return
}

/*
Where's your docs doc?!
*/
func testUnlessConditions(unless string, localCtx interface{}) (found interface{}, err error) {
	// HOFSTADTER_START testUnlessConditions

	// check the unless clause
	if unless != "" {
		logger.Info("Unless", "unless", unless)
		unlessElems, err := dotpath.Get(unless, localCtx, false)
		if err != nil {
			return nil, err
		}
		logger.Debug("  elems", "unlessElems", unlessElems)
		// nil err means something was found
		if unlessElems != nil {
			return unlessElems, nil
		}
	}

	// HOFSTADTER_END   testUnlessConditions
	return
}

/*
Where's your docs doc?!
*/
func testWhenConditions(when string, localCtx interface{}) (found interface{}, err error) {
	// HOFSTADTER_START testWhenConditions

	// check the when clause
	if when != "" {
		logger.Info("When", "when", when)

		// add AND (&&) and OR (||) here
		// make a loop to check, how to combine elements?
		var whenElems interface{}
		whenElems, err = dotpath.Get(when, localCtx, false)
		if err != nil {
			return
		}
		logger.Debug("  elems", "whenElems", whenElems)
		if whenElems == nil {
			return
		}
		switch W := whenElems.(type) {
		case []interface{}:
			// found, but empty, do we ever get here?
			if len(W) == 0 {
				logger.Warn("Skipping TemplatePair When Field: (empty array)", "when", when, "err", err, "whenElems", whenElems)
				return
			}
		}

		logger.Debug("When is NOW")
		return whenElems, nil
	}
	// HOFSTADTER_END   testWhenConditions
	return
}

// HOFSTADTER_BELOW
