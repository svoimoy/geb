package cmd

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	{{#each CLI.commands as |Cmd|}}
	{{#if Cmd.subcommands}}
	"{{{trimprefix file_fulldir (concat2 ENV.GOPATH '/src/')}}}/{{lower Cmd.name}}"
	{{/if}}
	{{/each}}
)

var logger = log.New()

func SetLogger(l log.Logger) {
	lcfg := viper.GetStringMap("log-config.cmd.default")

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

	// set subcommand loggers
	{{#each CLI.commands as |Cmd|}}
	{{lower Cmd.name}}.SetLogger(logger)
	{{/each}}
}