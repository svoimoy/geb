package run

import (
	// HOFSTADTER_START import
	"fmt"
	"os/exec"
	"strings"

	"github.com/hofstadter-io/geb/engine/project"
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
func Run(commands []string) (err error) {
	// HOFSTADTER_START Run
	proj := project.NewProject()

	err = proj.Load("geb.yaml", []string{})
	if err != nil {
		return errors.Wrapf(err, "While loading project")
	}

	for _, cname := range commands {
		Cfg, ok := proj.Config.RunConfig[cname]
		if ok {
			fmt.Printf("Exec'n %+v\n", Cfg)
		} else {
			return errors.Errorf("Unknown command %q", cname)
		}

		for _, C := range Cfg {
			cmdFlds := strings.Fields(C.Command)
			if len(cmdFlds) > 1 {
				C.Command = cmdFlds[0]
				C.Args = append(cmdFlds[1:], C.Args...)
			}

			cmd := exec.Command(C.Command, C.Args...)
			stdoutStderr, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error running command", C, err)
				fmt.Printf("%s\n", stdoutStderr)
				return err
			}
			fmt.Printf("%s\n", stdoutStderr)

		}
	}
	// HOFSTADTER_END   Run
	return
}

// HOFSTADTER_BELOW
