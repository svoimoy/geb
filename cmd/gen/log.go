package gen

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
args:
- help: Path to the input design file.
  name: designFile
  required: true
  type: string
- help: Path to the template file.
  name: templateFile
  required: true
  type: string
- help: Path to the output file. Can also be 'stdout'.
  name: outputFile
  required: true
  type: string
long: Generate a file from design and a template.
name: File
parent: Gen
path: commands.subcommands
short: Generate a file.
usage: file <designFile> <templateFile> <outputFile>

*/
