package engine

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/gen"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
	"github.ibm.com/hofstadter-io/geb/engine/project"
	"github.ibm.com/hofstadter-io/geb/engine/render"
	"github.ibm.com/hofstadter-io/geb/engine/system"
	"github.ibm.com/hofstadter-io/geb/engine/templates"
	"github.ibm.com/hofstadter-io/geb/engine/unify"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

var logger = log.New()

func SetLogger(l log.Logger) {

	lcfg := viper.GetStringMap("log-config.engine.default")

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

	design.SetLogger(logger)
	dsl.SetLogger(logger)
	gen.SetLogger(logger)
	plan.SetLogger(logger)
	project.SetLogger(logger)
	render.SetLogger(logger)
	system.SetLogger(logger)
	templates.SetLogger(logger)
	unify.SetLogger(logger)
	utils.SetLogger(logger)
}
