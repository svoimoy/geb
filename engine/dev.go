package engine

import (
	"fmt"

	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/project"
)

func Dev(args []string) {

	filename := "geb.yaml"
	if len(args) == 1 {
		filename = args[0]
	}

	proj := project.NewProject()

	fmt.Println("Reading config file from:", filename)
	c, err := project.ReadConfigFile(filename)
	if err != nil {
		fmt.Println("Error reading project config "+filename+":", err)
		return
	}
	proj.Config = c

	d_dir := proj.Config.DesignDir
	fmt.Println("Reading desings from folder:", d_dir)
	d, err := design.CreateFromFolder(d_dir)
	if err != nil {
		fmt.Println("Error reading project designs from "+d_dir+":", err)
		return
	}
	proj.Design = d
	// fmt.Printf("Project:\n%+v\n", proj)

	dcfg := proj.Config.DslConfig
	fmt.Println("DSL override order (first to last):")
	for _, path := range dcfg.Paths {
		fmt.Println("   ", path)
	}

	fmt.Println("Loading generators:")
	for _, blob := range dcfg.Default {
		fmt.Println("   ", blob)
		gen.NewGenerator()

	}

	dsl_dir := "dsl"
	fmt.Println("Searching for DSLs in:", dsl_dir)
	dsl.FindAvailable(dsl_dir)

	fmt.Println()
}
