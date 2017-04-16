package resources

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package resources

/*
ctx_path: dsl.lib.serve.api.resources.[0]
methods:
- method: list
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[0].output.[0]
    name: ts
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[0]/output
    pkgPath: serve/templates/ts
    type: array:lib.templates.template.views.short
- input:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[1].input.[0]
    name: in-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[1]/input
    pkgPath: serve/templates/in-tpl
    type: lib.templates.template.views.create
  method: post
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[1].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[1]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template
- method: get
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[2].output.[0]
    name: t
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[2]/output
    pkgPath: serve/templates/t
    type: lib.templates.template
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[2].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[2]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
- input:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].input.[0]
    name: in-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/input
    pkgPath: serve/templates/in-tpl
    type: lib.templates.template
  method: put
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[3].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[3]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
- method: delete
  output:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[4].output.[0]
    name: out-tpl
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[4]/output
    pkgPath: serve/templates/out-tpl
    type: lib.templates.template.views.short
  path-params:
  - ctx_path: dsl.lib.serve.api.resources.[0].methods.[4].path-params.[0]
    name: template-id
    parent: serve.templates
    parent_path: dsl.lib.serve.api.resources.[0]
    pkg_path: lib/serve/api/resources/[0]/methods/[4]/path-params
    pkgPath: serve/templates/template-id
    type: lib.templates.template.fields.id
name: templates
omit-db-calls: true
parent: serve
parent_path: dsl.lib.serve.api
path: resources
pkg_path: lib/serve/api/resources
pkgPath: serve/templates
resource: lib.templates.template
route: templates

*/

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config.resources.default")

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
	lcfg := viper.GetStringMap("log-config.resources")

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
