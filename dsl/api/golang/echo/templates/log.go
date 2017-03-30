{{#with RepeatedContext as |CTX| }}
{{#if (eq CTX.parent DslContext.name)}}
package {{camel CTX.path}}
{{else}}
package {{#if CTX.parent}}{{camel CTX.parent}}{{else}}commands{{/if}}
{{/if}}

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package {{#each (split CTX.pkg_path "/")}}{{#if @last }}{{camel .}}{{/if}}{{/each}}

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger = log.New()

func SetLogger(l log.Logger) {
	{{#if (eq CTX.parent DslContext.name)}}
		ldcfg := viper.GetStringMap("log-config.{{replace CTX.path '/' '.' -1 }}.default")
	{{else}}
		ldcfg := viper.GetStringMap("log-config.{{replace (trimto_first CTX.pkgPath '/' false) '/' '.' -1 }}.default")
	{{/if}}

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
	{{#each CTX.commands as |Cmd|}}
	{{#if Cmd.subcommands}}
	{{lower Cmd.name}}.SetLogger(logger)
	{{/if}}
	{{/each}}


	// possibly override locally
	{{#if (eq CTX.parent DslContext.name)}}
		lcfg := viper.GetStringMap("log-config.{{replace CTX.path '/' '.' -1 }}")
	{{else}}
		lcfg := viper.GetStringMap("log-config.{{replace (trimto_first CTX.pkgPath '/' false) '/' '.' -1 }}")
	{{/if}}

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

{{/with}}

