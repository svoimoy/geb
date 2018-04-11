package new

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"path/filepath"

	"github.com/hofstadter-io/geb/engine"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/plan"
	"github.com/hofstadter-io/geb/engine/project"
	"github.com/hofstadter-io/geb/engine/render"
	"github.com/hofstadter-io/geb/engine/utils"
	"github.com/pkg/errors"
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
func NewProject(data map[string]interface{}) (err error) {
	// HOFSTADTER_START NewProject
	name := data["name"].(string)
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	outfile := filepath.Join(outdir, "geb.yaml")
	err = engine.GenerateFileWithData(data, PACKAGE_TEMPLATE, outfile)
	if err != nil {
		return
	}

	return mkdirs(outdir, PACKAGE_DIRS)
	// HOFSTADTER_END   NewProject
	return
}

/*
Where's your docs doc?!
*/
func NewDsl(data map[string]interface{}) (err error) {
	// HOFSTADTER_START NewDsl
	name := data["name"].(string)
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	outfile := filepath.Join(outdir, "geb-dsl.yaml")
	err = engine.GenerateFileWithData(data, DSL_TEMPLATE, outfile)
	if err != nil {
		return
	}
	// HOFSTADTER_END   NewDsl
	return
}

/*
Where's your docs doc?!
*/
func NewGenerator(data map[string]interface{}) (err error) {
	// HOFSTADTER_START NewGenerator
	name := data["name"].(string)
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	// TODO search for geb-dsl.yaml recurisvely
	dsltype := "unknown"
	data["type"] = dsltype

	outfile := filepath.Join(outdir, "geb-gen.yaml")
	err = engine.GenerateFileWithData(data, GEN_TEMPLATE, outfile)
	if err != nil {
		return
	}
	return mkdirs(outdir, GEN_DIRS)
	// HOFSTADTER_END   NewGenerator
	return
}

/*
Where's your docs doc?!
*/
func NewDesign(data map[string]interface{}) (err error) {
	// HOFSTADTER_START NewDesign

	name := data["name"].(string)
	dslname := data["dslname"].(string)
	genname := data["genname"].(string)
	subname := data["subname"].(string)

	P := project.NewProject()

	// look for dsl
	file := utils.LookForKnownFiles()

	var G *gen.Generator
	switch file {
	case "geb.yml", "geb.yaml":
		err := P.Load(file, nil)
		if err != nil {
			return errors.Wrap(err, "in NewDesign\n")
		}

		found := false
		for key, D := range P.DslMap {
			if key == dslname {
				found = true
				var gOk bool
				G, gOk = D.Generators[genname]
				if !gOk {
					return errors.New("could not find the generator for dsl: " + dslname + " " + genname)
				}

				gfound := false
				for path, _ := range G.Config.NewConfigs {
					println("nf: ", path)
					if path == subname {
						gfound = true
						break
					}

				}
				if !gfound {
					return errors.New(dslname + "/" + genname + "/" + subname + " ...has no new files setup")
				}
				break
			}
		}

		if !found {
			return errors.New("could not find the dsl: " + dslname)
		}

	default:
		return errors.New("could not find a geb project file")
	}

	wd, err := os.Getwd()
	if err != nil {
		return
	}

	namepath := filepath.Dir(name)
	name = filepath.Base(name)

	filename := name + "-" + dslname + ".yaml"

	// TODO DesignDir is too specific
	outdir := filepath.Join(wd, P.Config.DesignDir, namepath)
	fmt.Println("ProjConfig", P.Config.DesignDir, outdir)
	/*
		err = os.MkdirAll(outdir, 0775)
		if err != nil {
			return
		}
	*/

	// TODO, process gen Newfiles
	outfile := filepath.Join(outdir, filename)

	fmt.Printf("new design: %q %v\n", outfile, G.Config.NewConfigs)

	plans, err := plan.MakeNewPlans(G, subname, outdir, data)
	if err != nil {
		fmt.Println("error planning new things", err)
		return
	}

	fmt.Println("Plans", plans)

	errs := render.RenderPlans(plans, outdir)
	if errs != nil && len(errs) > 0 {
		fmt.Println("Errors while rendering new stuff", errs)
		return errs[0]
	}

	// HOFSTADTER_END   NewDesign
	return
}

// HOFSTADTER_BELOW

func prep(inname string) (name, outdir string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	mkdir := true
	if inname == "" {
		mkdir = false
		name = filepath.Base(wd)
	} else {
		name = inname
	}

	if mkdir {
		outdir = filepath.Join(wd, name)
		err = os.Mkdir(outdir, 0775)
		if err != nil {
			return
		}
	} else {
		outdir = wd
	}

	return
}

func mkdirs(base string, dirs []string) (err error) {
	for _, dir := range dirs {
		d := filepath.Join(base, dir)
		err = os.MkdirAll(d, 0775)
		if err != nil {
			return
		}
	}

	return
}
