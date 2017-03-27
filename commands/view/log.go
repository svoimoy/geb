package view

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package subcommands

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"




)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.cli.commands.4.subcommands.default")
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
	lcfg := viper.GetStringMap("log-config.cli.commands.4.subcommands")

	if lcfg == nil || len(lcfg) == 0  {
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
aliases:
- P
args:
- ctx_path: dsl.cli.commands.4.subcommands.5.args.0
  help: one ore more dotpaths for indexing into the data
  name: paths
  parent: geb.view.plans
  parent_path: dsl.cli.commands.4.subcommands.5
  pkg_path: cli/commands/4/subcommands/5/args
  rest: true
  type: array:string
ctx_path: dsl.cli.commands.4.subcommands.5
long: View information about a Project's Plans known from the current path
name: plans
parent: view
parent_path: dsl.cli.commands.4
path: commands.subcommands
pkg_path: cli/commands/4/subcommands
short: View information about a Project's Plans
usage: plans <dotpath>...

*/


// HOFSTADTER_BELOW
