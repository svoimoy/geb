package gebberish

import (
	// HOFSTADTER_START import
	"strings"
	// HOFSTADTER_END   import

	"fmt"

	
	"github.com/spf13/cobra"

)

// Tool:   geb
// Name:   Mi
// Usage:  mi
// Parent: Gebberish
// ParentPath: 

var MiLong = `Welcome to the MI game

start with mi-string = 'MI'

mi-rule-1:    if mi-string ends in 'I',        you may add a 'U'
mi-rule-2:    suppose mi-string = 'Mx',          then you may make it 'Mxx'
mi-rule-3:    if mi-string contains an 'III',  you may replace it with 'U'
mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)

Goal: Try to get 'MU'

Input: r1,r2,r3,r4 or reset

Notes:
 - rules 3 and 4 require a zero-indexed postional argument to determine the occurance to remove
 - use 'reset' to restore mi-string to 'MI'
`





var MiCmd = &cobra.Command {
	Use: "mi",
	Short: "View information about a Project's Plans",
	Long: MiLong,
		
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("In MiCmd", "args", args)
		// Argument Parsing
		// [0]name:   rule
		//     help:   The rule to apply [r# or rule-#]
		//     req'd:  true
		if 0 >= len(args) {
			cmd.Usage()
			return
		}
		var rule string
		if 0 < len(args) {
			rule = args[0]
		}
		
		// [1]name:   extra
		//     help:   optional args to rules 3 and 4
		//     req'd:  
		var extra []string
			
		if 1 < len(args) {
			extra = args[1:]
		}
		
		

		// HOFSTADTER_START cmd_run
		rule = strings.Replace(rule, "-", "", 1)
		rule = strings.Replace(rule, "ule", "", 1)
		rule = strings.Replace(rule, "R", "r", 1)
		fmt.Printf("MI: (%s) [%v]\n", rule, extra)
		// HOFSTADTER_END   cmd_run
	},
		}


func init() {

}


/*
Repeated Context
----------------
args:
- help: The rule to apply [r# or rule-#]
  name: rule
  required: true
  type: string
- help: optional args to rules 3 and 4
  name: extra
  rest: true
  type: array:string
long: |
  Welcome to the MI game

  start with mi-string = 'MI'

  mi-rule-1:    if mi-string ends in 'I',        you may add a 'U'
  mi-rule-2:    suppose mi-string = 'Mx',          then you may make it 'Mxx'
  mi-rule-3:    if mi-string contains an 'III',  you may replace it with 'U'
  mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)

  Goal: Try to get 'MU'

  Input: r1,r2,r3,r4 or reset

  Notes:
   - rules 3 and 4 require a zero-indexed postional argument to determine the occurance to remove
   - use 'reset' to restore mi-string to 'MI'
name: Mi
parent: Gebberish
path: commands.subcommands
short: View information about a Project's Plans
usage: mi

*/
