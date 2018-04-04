package new

import (
	// HOFSTADTER_START import
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

	return mkdirs(outdir, []string{"design", "dsls"})
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

// HOFSTADTER_BELOW

const PACKAGE_TEMPLATE = `name: "__NAME__"
about: "All about __NAME__..."

output-dir: "."

design-dir: "design"

# Below is example content, update to your needs.
# See more at [doc link t.b.d.]

dsl-config:
  paths:
    - "$HOME/.geb/dsl"
  default:
    - dsl: cli
      gen:
        - golang
      output-dir: "."

build-pipeline:
  stages:
    - name: generate
      cmd: "geb gen"
    - name: format
      cmd: "gofmt -w {{config.output-dir}}"
    - name: build
      cmd: "go build -o {{config.output-dir}} {{config.output-dir}}"

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
