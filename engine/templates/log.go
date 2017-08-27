package templates

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.engine.templates.default")
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

	// set sub-loggers before possibly overriding locally next

	// HOFSTADTER_START logging-config
	// HOFSTADTER_END logging-config

	// possibly override locally
	lcfg := viper.GetStringMap("log-config.engine.templates")

	if lcfg == nil || len(lcfg) == 0 {
		logger = l
	} else {
		// hack because of default override (should look for both upfront)
		logger = log.New()

		// find the logging level
		level_iface, ok := lcfg["level"]
		level_str := "warn"
		if ok {
			level_str = level_iface.(string)
		}

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
