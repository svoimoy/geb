package engine

import (
	"github.com/hofstadter-io/geb/engine/design"
	"github.com/hofstadter-io/geb/engine/dsl"
	"github.com/hofstadter-io/geb/engine/gen"
	"github.com/hofstadter-io/geb/engine/project"
	"github.com/hofstadter-io/geb/engine/templates"
	"github.com/hofstadter-io/geb/engine/utils"
	log "gopkg.in/inconshreveable/log15.v2" // logging framework
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l

	design.SetLogger(l)
	dsl.SetLogger(l)
	gen.SetLogger(l)
	project.SetLogger(l)
	templates.SetLogger(l)
	utils.SetLogger(l)
}
