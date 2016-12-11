package cmd

import (
	// HOFSTADTER_START import
	"strings"

	"github.com/ryanuber/go-glob"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/project"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
	// HOFSTADTER_END   import

	"fmt"

	"github.com/spf13/cobra"
)

// Tool:   geb
// Name:   Create
// Usage:  create <name> <dsl/gen>...
// Parent: geb

var CreateLong = `Create a new project with the given name.
Optionally specifiy the starting set of
DSLs and generators for the project.
The output directory defaults to the same name,
unless overridden by the output flag.
`

var CreateCmd = &cobra.Command{
	Use: "create <name> <dsl/gen>...",
	Aliases: []string{
		"new",
	},
	Short: "Create a new project",
	Long:  CreateLong,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In CreateCmd", "args", args)
		// Argument Parsing
		// [0]name:   name
		//     help:   The name of the new project to create
		//     req'd:  true
		if 0 >= len(args) {
			cmd.Usage()
			return
		}
		var name string
		if 0 < len(args) {
			name = args[0]
		}

		// [1]name:   dsls_n_gens
		//     help:   The starting list of DSLs and generators by path.
		//     req'd:
		var dsls_n_gens []string

		if 1 < len(args) {
			dsls_n_gens = args[1:]
		}

		// HOFSTADTER_START cmd_run
		tpl_data := map[string]interface{}{}
		tpl_data["name"] = name
		tpl_data["dsls"] = map[string]map[string]interface{}{}

		rpath, err := utils.ResolvePath("$HOME/.hofstadter/dsl")
		if err != nil {
			fmt.Println("Error:", err)
		}

		avail, err := dsl.FindAvailable(rpath)
		if err != nil {
			fmt.Println("Error:", err)
		}

		for _, d := range dsls_n_gens {
			flds := strings.Split(d, "/")
			fmt.Println("Fields:", flds)
			found := false
			for i, _ := range flds {
				dsl_path := strings.Join(flds[:i+1], "/")
				fmt.Println("  dsl-path:", dsl_path)
				d_dsl, d_ok := avail[dsl_path]
				if d_ok {
					gen_path := "*"
					spath := gen_path
					if i < len(flds) {
						gen_path = strings.Join(flds[i+1:], "/")
						spath = gen_path + "*"
					}
					fmt.Println("  spath:", spath)
					for _, path := range d_dsl.AvailableGenerators {
						found = glob.Glob(spath, path)
						logger.Debug("GLOB:", "spath", spath, "path", path, "found", found)
						if found {
							// add to the tmp_data
							d_dsls := tpl_data["dsls"].(map[string]map[string]interface{})
							dsl_data, ok := d_dsls[dsl_path]
							if !ok {
								dsl_data = map[string]interface{}{}
								dsl_data["name"] = dsl_path
								d_dsls[dsl_path] = dsl_data
							}
							if gen_path != "*" && gen_path != "" {
								if _, ok := dsl_data["gens"]; !ok {
									dsl_data["gens"] = map[string]string{}
								}
								d_gens := dsl_data["gens"].(map[string]string)
								_, ok = d_gens[gen_path]
								if !ok {
									d_gens[gen_path] = gen_path
								}
							}

							break
						}
					} // end for loop looking for gen in available generators
				}
				if found {
					break
				}
			}
			if !found {
				fmt.Println("failed to fine dsl/gen:", d)
				return
			}
		}

		P := project.New()
		P.Config = project.NewConfig()
		P.Config.OutputDir = name
		P.Config.DslConfig.Paths = []string{"$HOME/.hofstadter/dsl"}
		default_proj := project.GenPair{
			Dsl:       "project/fresh",
			Gen:       []string{"default"},
			OutputDir: ".",
		}
		P.Config.DslConfig.Default = []project.GenPair{default_proj}

		err = P.FindAvailableGenerators([]string{"$HOME/.hofstadter/dsl"})
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = P.LoadDefaultGenerators()
		if err != nil {
			fmt.Println("Error:", err)
		}

		P.Design.Dsl["project"] = tpl_data

		err = P.Plan()
		if err != nil {
			fmt.Println("Error:", err)
		}

		err = P.Render()
		if err != nil {
			fmt.Println("Error:", err)
		}

		// HOFSTADTER_END   cmd_run
	},
}

func init() {
	RootCmd.AddCommand(CreateCmd)

}

/*
Repeated Context
----------------
aliases:
- new
args:
- help: The name of the new project to create
  name: name
  required: true
  type: string
- help: The starting list of DSLs and generators by path.
  name: dsls_n_gens
  rest: true
  type: array:string
long: |
  Create a new project with the given name.
  Optionally specifiy the starting set of
  DSLs and generators for the project.
  The output directory defaults to the same name,
  unless overridden by the output flag.
name: Create
parent: geb
path: commands
short: Create a new project
usage: create <name> <dsl/gen>...

*/
