package commands

// The following line in the template needs fixing, it's probably related to the tree traversal and adding information
// go unification improvements!!
// package 

import (
	"github.com/spf13/viper"
	log "gopkg.in/inconshreveable/log15.v2"

		"github.ibm.com/hofstadter-io/geb/commands/gebberish"
			"github.ibm.com/hofstadter-io/geb/commands/system"
		"github.ibm.com/hofstadter-io/geb/commands/view"
		"github.ibm.com/hofstadter-io/geb/commands/gen"
		


)

var logger = log.New()

func SetLogger(l log.Logger) {
	ldcfg := viper.GetStringMap("log-config..default")
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
			system.SetLogger(logger)
		view.SetLogger(logger)
		gen.SetLogger(logger)
		



	// possibly override locally
	lcfg := viper.GetStringMap("log-config.")

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

/*
commands:
- aliases:
  - new
  args:
  - ctx_path: dsl.cli.commands.0.args.0
    help: The name of the new project to create
    name: name
    parent: geb.create
    parent_path: dsl.cli.commands.0
    pkg_path: cli/commands/0/args
    required: true
    type: string
  - ctx_path: dsl.cli.commands.0.args.1
    help: The starting list of DSLs and generators by path.
    name: dsls_n_gens
    parent: geb.create
    parent_path: dsl.cli.commands.0
    pkg_path: cli/commands/0/args
    rest: true
    type: array:string
  ctx_path: dsl.cli.commands.0
  long: |
    Create a new project with the given name.
    Optionally specifiy the starting set of
    DSLs and generators for the project.
    The output directory defaults to the same name,
    unless overridden by the output flag.
  name: create
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: Create a new project
  usage: create <name> <dsl/gen>...
- aliases:
  - games
  - G
  ctx_path: dsl.cli.commands.1
  hidden: true
  long: Games, shenanigans, and other gebberish.
  name: gebberish
  omit-run: true
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: it's a puzzle?!
  subcommands:
  - args:
    - ctx_path: dsl.cli.commands.1.subcommands.0.args.0
      help: The rule to apply [r# or rule-#]
      name: rule
      parent: geb.gebberish.mi
      parent_path: dsl.cli.commands.1.subcommands.0
      pkg_path: cli/commands/1/subcommands/0/args
      required: true
      type: string
    - ctx_path: dsl.cli.commands.1.subcommands.0.args.1
      help: optional args to rules 3 and 4
      name: extra
      parent: geb.gebberish.mi
      parent_path: dsl.cli.commands.1.subcommands.0
      pkg_path: cli/commands/1/subcommands/0/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.1.subcommands.0
    long: |
      Welcome to the MI game

      start with mi-string = 'MI'

      mi-rule-1:    if mi-string ends in 'I',        you may add a 'U'
      mi-rule-2:    suppose mi-string = 'Mx',          then you may make it 'Mxx'
      mi-rule-3:    if mi-string contains an 'III',  you may replace it with 'U'
      mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)

      Goal: Try to get 'MU'

      Input:
        - rules, h, help
        - c, curr, current, s, stat, status, get
        - reset, give-up, giveup, start-over, startover
        - 1, r1, rule1, rule-1
        - 2, r2, rule2, rule-2
        - 3, r3, rule3, rule-3 [pos]  (default is last pos)
        - 4, r4, rule4, rule-4 [pos]  (default is last pos)
    name: mi
    parent: gebberish
    parent_path: dsl.cli.commands.1
    path: commands.subcommands
    pkg_path: cli/commands/1/subcommands
    short: View information about a Project's Plans
    usage: mi
  usage: gebberish
- aliases:
  - i
  - convert
  - eat
  args:
  - ctx_path: dsl.cli.commands.2.args.0
    help: Path to the file or folder. Can also be 'stdin'.
    name: input
    parent: geb.import
    parent_path: dsl.cli.commands.2
    pkg_path: cli/commands/2/args
    required: true
    type: string
  - ctx_path: dsl.cli.commands.2.args.1
    help: Path to the output file or folder. Can also be 'stdout'.
    name: output
    parent: geb.import
    parent_path: dsl.cli.commands.2
    pkg_path: cli/commands/2/args
    required: true
    type: string
  ctx_path: dsl.cli.commands.2
  flags:
  - ctx_path: dsl.cli.commands.2.flags.0
    help: The type of input data to force geb to use a certain format
    long: type
    name: Type
    parent: geb.import
    parent_path: dsl.cli.commands.2
    pkg_path: cli/commands/2/flags
    short: T
    type: string
  long: |
    Import other stuff into Hofstadter.

    Stuff is...
      - json/jsonl, yaml, xml, protobuf, taml
      - swagger, goa.design
      - golang type definitions
  name: import
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: Import other stuff into Hofstadter
  usage: import <file or directory> <output file or directory>
- aliases:
  - sys
  - s
  ctx_path: dsl.cli.commands.3
  long: Manage the geb system and congiuration
  name: system
  omit-run: true
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: Manage the geb system and congiuration
  subcommands:
  - aliases:
    - initialize
    - setup
    ctx_path: dsl.cli.commands.3.subcommands.0
    long: Intializes the geb tool and the ~/.hofstadter dot folder.
    name: init
    parent: system
    parent_path: dsl.cli.commands.3
    path: commands.subcommands
    pkg_path: cli/commands/3/subcommands
    short: Initialize the geb tool and files
    usage: init
  - ctx_path: dsl.cli.commands.3.subcommands.1
    long: Copy development files to the dot folder
    name: dev-copy-in
    parent: system
    parent_path: dsl.cli.commands.3
    path: commands.subcommands
    pkg_path: cli/commands/3/subcommands
    short: Copy development files to the dot folder
    usage: dev-copy-in
  - ctx_path: dsl.cli.commands.3.subcommands.2
    long: Update the geb library DSLs, designs, and other files in the dot folder.
    name: update
    parent: system
    parent_path: dsl.cli.commands.3
    path: commands.subcommands
    pkg_path: cli/commands/3/subcommands
    short: Update the geb library and dot folder
    usage: update
  usage: system
- aliases:
  - v
  ctx_path: dsl.cli.commands.4
  long: View information known to the geb tool.
  name: view
  omit-run: true
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: View information known to the geb tool.
  subcommands:
  - aliases:
    - s
    - system
    - geb
    - config
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.0.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.system
      parent_path: dsl.cli.commands.4.subcommands.0
      pkg_path: cli/commands/4/subcommands/0/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.0
    long: View information about the global geb config
    name: system
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about Global geb config
    usage: sys <dotpaths>...
  - aliases:
    - d
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.1.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.dsl
      parent_path: dsl.cli.commands.4.subcommands.1
      pkg_path: cli/commands/4/subcommands/1/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.1
    long: View information about DSLs known from the current path
    name: dsl
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about DSLs
    usage: dsl <dotpath>...
  - aliases:
    - g
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.2.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.gen
      parent_path: dsl.cli.commands.4.subcommands.2
      pkg_path: cli/commands/4/subcommands/2/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.2
    long: View information about generators known from the current path
    name: gen
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about Generators
    usage: gen <dotpaths>...
  - aliases:
    - p
    - proj
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.3.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.project
      parent_path: dsl.cli.commands.4.subcommands.3
      pkg_path: cli/commands/4/subcommands/3/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.3
    long: View information about a Project known from the current path
    name: project
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about a Project
    usage: project <dotpath>...
  - aliases:
    - D
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.4.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.design
      parent_path: dsl.cli.commands.4.subcommands.4
      pkg_path: cli/commands/4/subcommands/4/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.4
    long: View information about Designs known from the current path
    name: design
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about Designs
    usage: design <dotpath>...
  - aliases:
    - P
    args:
    - ctx_path: dsl.cli.commands.4.subcommands.5.args.0
      help: one ore more dotpaths for indexing into the data
      name: paths
      parent: geb.view.plans
      parent_path: dsl.cli.commands.4.subcommands.5
      pkg_path: cli/commands/4/subcommands/5/args
      rest: true
      type: array:string
    ctx_path: dsl.cli.commands.4.subcommands.5
    long: View information about a Project's Plans known from the current path
    name: plans
    parent: view
    parent_path: dsl.cli.commands.4
    path: commands.subcommands
    pkg_path: cli/commands/4/subcommands
    short: View information about a Project's Plans
    usage: plans <dotpath>...
  usage: view
- aliases:
  - geb
  - geberate
  - generate
  - g
  ctx_path: dsl.cli.commands.5
  long: Generate a project from its working directory.
  name: gen
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: Generate a project.
  subcommands:
  - args:
    - ctx_path: dsl.cli.commands.5.subcommands.0.args.0
      help: Path to the input design file.
      name: designFile
      parent: geb.gen.file
      parent_path: dsl.cli.commands.5.subcommands.0
      pkg_path: cli/commands/5/subcommands/0/args
      required: true
      type: string
    - ctx_path: dsl.cli.commands.5.subcommands.0.args.1
      help: Path to the template file.
      name: templateFile
      parent: geb.gen.file
      parent_path: dsl.cli.commands.5.subcommands.0
      pkg_path: cli/commands/5/subcommands/0/args
      required: true
      type: string
    - ctx_path: dsl.cli.commands.5.subcommands.0.args.2
      help: Path to the output file. Can also be 'stdout'.
      name: outputFile
      parent: geb.gen.file
      parent_path: dsl.cli.commands.5.subcommands.0
      pkg_path: cli/commands/5/subcommands/0/args
      required: true
      type: string
    ctx_path: dsl.cli.commands.5.subcommands.0
    long: Generate a file from design and a template.
    name: file
    parent: gen
    parent_path: dsl.cli.commands.5
    path: commands.subcommands
    pkg_path: cli/commands/5/subcommands
    short: Generate a file.
    usage: file <designFile> <templateFile> <outputFile>
  usage: gen
- aliases:
  - b
  args:
  - ctx_path: dsl.cli.commands.6.args.0
    help: The stages to run in order. Used to override the pipeline in the project
      file.
    name: stages
    parent: geb.build
    parent_path: dsl.cli.commands.6
    pkg_path: cli/commands/6/args
    rest: true
    type: array:string
  ctx_path: dsl.cli.commands.6
  long: |
    Run the build pipeline specified in your project.
    Use this to run pre and post steps around 'gen gen'.
    Pipelines are also used by generators.
    See [...] for more information.
  name: build
  parent: geb
  parent_path: dsl.cli
  path: commands
  pkg_path: cli/commands
  short: Run the build pipeline for a project.
  usage: build
ctx_path: dsl.cli
long: |
  Hofstadter is a Framework
  for building data-centric
  Platforms. geb is the tool.
name: geb
omit-run: true
parent: ""
parent_path: ""
pflags:
- ctx_path: dsl.cli.pflags.0
  default: geb.yaml
  help: A geb project config file.
  long: config
  name: config
  parent: geb
  parent_path: dsl.cli
  pkg_path: cli/pflags
  short: c
  type: string
- ctx_path: dsl.cli.pflags.1
  default: design
  help: The design files directory.
  long: design
  name: design
  parent: geb
  parent_path: dsl.cli
  pkg_path: cli/pflags
  short: d
  type: string
- ctx_path: dsl.cli.pflags.2
  default: templates:~/.hofstadter/templates
  help: The search path for templates, reads from left to right, overriding along
    the way.
  long: template-paths
  name: template-paths
  parent: geb
  parent_path: dsl.cli
  pkg_path: cli/pflags
  short: t
  type: string
- ctx_path: dsl.cli.pflags.3
  default: output
  help: The directory to output generated files to.
  long: output
  name: output
  parent: geb
  parent_path: dsl.cli
  pkg_path: cli/pflags
  short: o
  type: string
pkg_path: ""
short: geb is the Hofstadter framework CLI tool

*/


// HOFSTADTER_BELOW
