package view

import (
	log "gopkg.in/inconshreveable/log15.v2"
)

var logger log.Logger

func SetLogger(l log.Logger) {
	logger = l
}

/*
Repeated Context
----------------
aliases:
- P
args:
- help: one ore more dotpaths for indexing into the data
  name: paths
  rest: true
  type: array:string
long: View information about a Project's Plans known from the current path
name: Plans
parent: View
path: commands.subcommands
short: View information about a Project's Plans
usage: plans <dotpath>...

*/
