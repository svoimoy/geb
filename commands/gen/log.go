package gen

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package subcommands

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"




)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.cli.commands.5.subcommands.default")
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
	lcfg := viper.GetStringMap("log-config.cli.commands.5.subcommands")

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
args:
- ctx_path: dsl.cli.commands.5.subcommands.0.args.0
  help: Path to the input design file.
  name: designFile
  parent: geb.gen.file
  parent_path: dsl.cli.commands.5.subcommands.0
  pkg_path: cli/commands/5/subcommands/0/args
  required: true
  type: string
- ctx_path: dsl.cli.commands.5.subcommands.0.args.1
  help: Path to the template file.
  name: templateFile
  parent: geb.gen.file
  parent_path: dsl.cli.commands.5.subcommands.0
  pkg_path: cli/commands/5/subcommands/0/args
  required: true
  type: string
- ctx_path: dsl.cli.commands.5.subcommands.0.args.2
  help: Path to the output file. Can also be 'stdout'.
  name: outputFile
  parent: geb.gen.file
  parent_path: dsl.cli.commands.5.subcommands.0
  pkg_path: cli/commands/5/subcommands/0/args
  required: true
  type: string
ctx_path: dsl.cli.commands.5.subcommands.0
long: Generate a file from design and a template.
name: file
parent: gen
parent_path: dsl.cli.commands.5
path: commands.subcommands
pkg_path: cli/commands/5/subcommands
short: Generate a file.
usage: file <designFile> <templateFile> <outputFile>

*/


// HOFSTADTER_BELOW
