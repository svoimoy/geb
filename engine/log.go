package engine

import (
	"github.ibm.com/hofstadter-io/geb/engine/design"
	"github.ibm.com/hofstadter-io/geb/engine/dsl"
	"github.ibm.com/hofstadter-io/geb/engine/gen"
	"github.ibm.com/hofstadter-io/geb/engine/plan"
	"github.ibm.com/hofstadter-io/geb/engine/project"
	"github.ibm.com/hofstadter-io/geb/engine/render"
	"github.ibm.com/hofstadter-io/geb/engine/templates"
	"github.ibm.com/hofstadter-io/geb/engine/utils"
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l

	design.SetLogger(l)
	dsl.SetLogger(l)
	gen.SetLogger(l)
	plan.SetLogger(l)
	project.SetLogger(l)
	render.SetLogger(l)
	templates.SetLogger(l)
	utils.SetLogger(l)
}
