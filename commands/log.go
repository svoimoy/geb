package commands

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.ibm.com/hofstadter-io/geb/commands/gebberish"
	"github.ibm.com/hofstadter-io/geb/commands/gen"
	"github.ibm.com/hofstadter-io/geb/commands/system"
	"github.ibm.com/hofstadter-io/geb/commands/view"
)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.commands.default")
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
	gebberish.SetLogger(logger)
	gen.SetLogger(logger)
	view.SetLogger(logger)
	system.SetLogger(logger)

	// possibly override locally
	lcfg := viper.GetStringMap("log-config.commands")

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

// HOFSTADTER_BELOW
