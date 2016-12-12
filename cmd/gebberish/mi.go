package gebberish

import (
	// HOFSTADTER_START import
	"strconv"

	"github.ibm.com/hofstadter-io/geb/gebberish"
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

Input:
  - rules, h, help
  - c, curr, current, s, stat, status, get
  - reset, give-up, giveup, start-over, startover
  - 1, r1, rule1, rule-1
  - 2, r2, rule2, rule-2
  - 3, r3, rule3, rule-3 [pos]  (default is last pos)
  - 4, r4, rule4, rule-4 [pos]  (default is last pos)
`

var MiCmd = &cobra.Command{
	Use:   "mi",
	Short: "View information about a Project's Plans",
	Long:  MiLong,

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
		switch rule {
		case "rules", "help", "h":
			fmt.Println(MiLong)
			return
		}
		i_arg := -1
		if len(extra) > 0 {
			i, err := strconv.Atoi(extra[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			i_arg = i
		}
		MI, err := gebberish.Mi(rule, i_arg)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("\nMI = %q\n\n", MI)

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
