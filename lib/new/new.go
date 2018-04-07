package new

import (
	// HOFSTADTER_START import
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
func NewProject(name string) (err error) {
	// HOFSTADTER_START NewProject
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	outfile := filepath.Join(outdir, "geb.yaml")
	output := strings.Replace(PACKAGE_TEMPLATE, "__NAME__", name, -1)
	err = ioutil.WriteFile(outfile, []byte(output), 0664)
	if err != nil {
		return
	}

	return mkdirs(outdir, []string{"design", "dsl", ".geb/shadow", ".geb/tmp"})
	// HOFSTADTER_END   NewProject
	return
}

/*
Where's your docs doc?!
*/
func NewDsl(name string) (err error) {
	// HOFSTADTER_START NewDsl
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	outfile := filepath.Join(outdir, "geb-dsl.yaml")
	output := strings.Replace(DSL_TEMPLATE, "__NAME__", name, -1)
	err = ioutil.WriteFile(outfile, []byte(output), 0664)
	if err != nil {
		return
	}
	// HOFSTADTER_END   NewDsl
	return
}

/*
Where's your docs doc?!
*/
func NewGenerator(name string) (err error) {
	// HOFSTADTER_START NewGenerator
	var outdir string
	name, outdir, err = prep(name)
	if err != nil {
		return
	}

	dsltype := "unknown"
	// TODO search for geb-dsl.yaml recurisvely

	outfile := filepath.Join(outdir, "geb-gen.yaml")
	output := strings.Replace(GEN_TEMPLATE, "__NAME__", name, -1)
	output = strings.Replace(output, "__TYPE__", dsltype, -1)
	err = ioutil.WriteFile(outfile, []byte(output), 0664)
	if err != nil {
		return
	}
	return mkdirs(outdir, []string{"designs", "partials", "templates", "new"})
	// HOFSTADTER_END   NewGenerator
	return
}

/*
Where's your docs doc?!
*/
func NewDesign(dslname string, genname string, name string) (err error) {
	// HOFSTADTER_START NewDesign

	subname := "default"
	if strings.Contains(genname, ":") {
		flds := strings.Split(genname, ":")
		genname, subname = flds[0], flds[1]
	}

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

	data := map[string]interface{}{
		"name": name,
		"dsl":  dslname,
		"gen":  genname,
		"sub":  subname,
	}

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

const PACKAGE_TEMPLATE = `name: "__NAME__"
about: "All about __NAME__..."

output-dir: "."

design-dir: "design"

# Below is example content, update to your needs.
# See more at [doc link t.b.d.]

dsl-config:
  paths:
    - "./dsl"
  default:
    - dsl: cli
      gen:
        - golang
      output-dir: "."

run-config:
  all:
    - name: "all"
      command: "geb run"
      args:
        - gen
        - fmt
        - install
  regen:
    - name: "regen"
      command: "geb run"
      args:
        - gen
        - fmt
  gen:
    - name: generate
      command: "geb"
      args:
        - "gen"
  fmt:
    - name: format
      command: "gofmt -w main.go commands engine lib"
  build:
    - name: build
      command: "go build"
  install:
    - name: install
      command: "go install"


log-config:
  default:
    level: warn
    stack: false
`

const DSL_TEMPLATE = `name: "__NAME__"
about: "All about __NAME__..."
version: "0.0.1"
type: "dsl"
`
const GEN_TEMPLATE = `name: "__NAME__"
about: "All about __NAME__..."
version: "0.0.1"
language: golang
type: "__TYPE__"

# Below is example content, update to your needs.
# See more at [doc link t.b.d.]

dependencies:
  generators:
    - dsl: common
      gen:
        - golang
      output-dir: "."
    - dsl: type
      gen:
        - golang
      output-dir: "."
    - dsl: pkg
      gen:
        - golang
      output-dir: "."

template-configs:

  - name: once-files
    field: "."
    templates:
      - in: "main.go"
        out: "main.go"
      - in: "root.go"
        out: "commands/root.go"
      - unless: commands
        in: "rootonlylog.go"
        out: "commands/log.go"
      - when: commands
        in: "log.go"
        out: "commands/log.go"

  - name: command
    field: commands
    templates:
      - in: "cmd.go"
        out: "commands/{{trimto_first pkgPath '/' false}}.go"
      - when: commands
        in: "log.go"
        out: "commands/{{name}}/log.go"
`

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
