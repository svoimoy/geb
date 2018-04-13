package run

import (
	// HOFSTADTER_START import
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"

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
		var pargs []string
		pipeline := cname
		if strings.ContainsAny(cname, "[]") {
			lhs := strings.Index(cname, "[")
			rhs := strings.Index(cname, "]")
			pipeline = cname[:lhs]
			pargstr := cname[lhs+1 : rhs]
			if strings.ContainsAny(pargstr, ",;") {
				pargs = strings.FieldsFunc(pargstr, func(r rune) bool {
					if r == ',' || r == ';' || unicode.IsSpace(r) {
						return true
					}
					return false
				})
			} else {
				pargs = strings.Fields(pargstr)
			}
		}
		Cfg, ok := proj.Config.RunConfig[pipeline]
		if ok {
			// fmt.Printf("Exec'n pipeline %q with %v\n%+v\n", pipeline, pargs, Cfg)
		} else {
			return errors.Errorf("Unknown command %q %v", pipeline, pargs)
		}

		for _, C := range Cfg {
			cmdFlds := strings.Fields(C.Command)
			if len(cmdFlds) > 1 {
				C.Command = cmdFlds[0]
				C.Args = append(cmdFlds[1:], C.Args...)
			}

			tmpArgs := C.Args
			C.Args = []string{}
			for _, arg := range tmpArgs {
				// splice in pargs
				if arg == "{args}" {
					C.Args = append(C.Args, pargs...)

				} else if len(arg) > 5 && arg[:5] == "{args[" {
					// parse index
					/*
						if strings.Contains(istr, ":") {
							iargstr = strings.Split(pargstr, ":")
							var i1, i2 int
							var err error

							index1, err = strconv.Atoi(istr)
							if err == nil {
								return errors.Wrapf(err, "in Run:cname[%v]", pargs)
							}
							if istr != ""
							index2, err = strconv.Atoi(istr)
							if err == nil {
								return errors.Wrapf(err, "in Run:cname[%v]", pargs)
							}
							if index1 == -1
							if index1 >= 0 && index1 < index2 && index2 < len(pargs) {
								C.Args = append(C.Args, pargs[index])
							}
						} else {
							index, err := strconv.Atoi(istr)
							if err == nil {
								return errors.Wrapf(err, "in Run:cname[%v]", pargs)
							}

							if index >= 0 && index < len(pargs) {
								C.Args = append(C.Args, pargs[index])
							} else {
								return errors.Errorf("index out of bounds in %q with %v")
							}
						}
					*/
				} else {
					C.Args = append(C.Args, arg)
				}

			}

			cmd := exec.Command(C.Command, C.Args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error running command", C, err)
				return err
			}

		}
	}
	// HOFSTADTER_END   Run
	return
}

// HOFSTADTER_BELOW
