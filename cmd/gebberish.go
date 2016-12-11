package cmd

import (
	// HOFSTADTER_START import
	// HOFSTADTER_END   import


	
	"github.com/spf13/cobra"

	"github.ibm.com/hofstadter-io/geb/cmd/gebberish"
)

// Tool:   geb
// Name:   Gebberish
// Usage:  gebberish
// Parent: geb

var GebberishLong = `Games, shenanigans, and other gebberish.`





var GebberishCmd = &cobra.Command {
	Hidden: true,
	Use: "gebberish",
	Aliases: []string{ 
		"games",
"G",
	},
	Short: "it's a puzzle?!",
	Long: GebberishLong,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		gebberish.SetLogger(logger)
		logger.Debug("In PersistentPreRun GebberishCmd", "args", args)

		// HOFSTADTER_START cmd_persistent_prerun
		// HOFSTADTER_END   cmd_persistent_prerun
	},
			}


func init() {
	RootCmd.AddCommand(GebberishCmd)

	GebberishCmd.AddCommand(gebberish.MiCmd)

	gebberish.SetLogger(logger)
}


/*
Repeated Context
----------------
aliases:
- games
- G
hidden: true
long: Games, shenanigans, and other gebberish.
name: Gebberish
omit-run: true
parent: geb
path: commands
short: it's a puzzle?!
subcommands:
- args:
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
usage: gebberish

*/
