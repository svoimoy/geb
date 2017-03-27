package gebberish

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package subcommands

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.cli.commands.1.subcommands.default")
	if ldcfg == nil || len(ldcfg) == 0 {
		logger = l
	} else {
		// find the logging level
		level_str := ldcfg["level"].(string)
		level, err := log.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}

		// possibly find the stack switch
		stack := false
		stack_tmp := ldcfg["stack"]
		if stack_tmp != nil {
			stack = stack_tmp.(bool)
		}

		// build the local logger
		termlog := log.LvlFilterHandler(level, log.StdoutHandler)
		if stack {
			term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
			termlog = log.LvlFilterHandler(level, term_stack)
		}

		// set the local logger
		logger.SetHandler(termlog)
	}

	// set subcommand loggers before possibly overriding locally next

	// possibly override locally
	lcfg := viper.GetStringMap("log-config.cli.commands.1.subcommands")

	if lcfg == nil || len(lcfg) == 0 {
		logger = l
	} else {
		// find the logging level
		level_str := lcfg["level"].(string)
		level, err := log.LvlFromString(level_str)
		if err != nil {
			panic(err)
		}

		// possibly find the stack switch
		stack := false
		stack_tmp := lcfg["stack"]
		if stack_tmp != nil {
			stack = stack_tmp.(bool)
		}

		// build the local logger
		termlog := log.LvlFilterHandler(level, log.StdoutHandler)
		if stack {
			term_stack := log.CallerStackHandler("%+v", log.StdoutHandler)
			termlog = log.LvlFilterHandler(level, term_stack)
		}

		// set the local logger
		logger.SetHandler(termlog)
	}

}

/*
args:
- ctx_path: dsl.cli.commands.1.subcommands.0.args.0
  help: The rule to apply [r# or rule-#]
  name: rule
  parent: geb.gebberish.mi
  parent_path: dsl.cli.commands.1.subcommands.0
  pkg_path: cli/commands/1/subcommands/0/args
  required: true
  type: string
- ctx_path: dsl.cli.commands.1.subcommands.0.args.1
  help: optional args to rules 3 and 4
  name: extra
  parent: geb.gebberish.mi
  parent_path: dsl.cli.commands.1.subcommands.0
  pkg_path: cli/commands/1/subcommands/0/args
  rest: true
  type: array:string
ctx_path: dsl.cli.commands.1.subcommands.0
long: |
  Welcome to the MI game

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
name: mi
parent: gebberish
parent_path: dsl.cli.commands.1
path: commands.subcommands
pkg_path: cli/commands/1/subcommands
short: View information about a Project's Plans
usage: mi

*/

// HOFSTADTER_BELOW
