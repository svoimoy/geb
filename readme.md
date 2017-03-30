# geb: data-centric design, output-agnostic creation

__geb__ is the command-line tool for the hofstadter
framework for data-centric design and output-agnostic creation
of just about anything.

With it, you can:

- Create designs and generate code in multiple languages in parallel
- Modify code and designs, regenerate without disruption
- Customize and extend templates, generators, and dsls.
- Share designs and templates across your applications and organization.
- Contribute to the ecosystem by sharing your designs, templates, generators, or complete applications.

### Installation

1. Clone this repository to: `$GOPATH/github.ibm.com/hofstadter-io/geb`
2. Change to that directory and run: `go install`

It may also be `go get`-able, just don't use `-u` because the dependencies
are vendored in and go gets confused. (not extensively tested)

__Then do the initial setup:__

See the Tips section on getting setup for now.
You can copy rather than link if you do not plan to develop geb.


### Usage

The docs are pretty much m.i.a. but you should still dive in!

Once you've followed the tips below,
start by copying a design folder and geb.yaml
from an example and running `geb gen`.

The basic layout for a project would be:

```
project
  - geb.yaml
  - design/
    - api.yaml
    - user.yaml
    - post.yaml
    - db.yaml
```

Many of the examples are out of data right now.
This repository and [Xtalk](https://github.ibm.com/krobots/k8s-csf-xtalk)
are both up-to-date examples. (minus Xtalk's pkg type usage in function designs)



### Developing

The docs are pretty much m.i.a. but you should still dive in!


##### Overriding generators

To override a generator, in any project directory:

Setup:

1. make a dir `dsl`
1. add sub-dirs to `dsl`
  1. match the directory structure of the library dsls
  1. Be sure to include the `geb-dsl.yaml` and `geb-gen.yaml` files
1. make sure the `dsl-config.paths` has the correct entries.

To __override__, add the file
to the expected location.

To __extend__ the templates,
add any new files you wish to.
Repeated templates also need
to be associated in the `geb-gen.yaml` file.

##### Extending dsls

DSLs are also easily extended.
Since both dsl and design are data
both are quite flexible together,
and combine for simple extensibility
and customization
(with the ability to override generators as well).

To extend a dsl:

1. Add any additional data to the design file, in the dsl specs.
1. Reference the new fields in your overriden generators.

Eventually DSLs will become more formal
and the documentation will also
become more detailed on the
dsl specifics (types too)
 



### Tips

Some setup steps and tips while things are still shakey.

Make the following symlinks,
from the development folder:

1. the `geb` binary to `/usr/local/bin/geb`
1. the `dsl` folder go `~/.geb/dsl`

This will keep you defaulting to
the latest build and templates,
when working in any directory
and letting you override
locally while still developing.



### Documentation

Coming soon:

- Getting started
  - Installing
  - Generating your first file
  - Overview

- Concepts
  - Data centric design
  - Designs
  - DSLs
  - Generators
  - Projects
  - Merging, layering, overriding

- Project Walkthrough
  - Starting an API
  - Adding types
  - Adding routes and resources
  - The `geb.yaml` project file
    - main configuration
    - generators
    - build pipeline 

- Designs
  - Overview
  - Organization
  - Output
  - Main DSLs
    - types
    - pkg
    - api
    - cli

- Exploring the DSLs
  - common
  - configuration
  - validation
  - testing
  - databases
  - binding and ETL
  - calls, sequences, map-reduce, parallel

- Additional Topics
  - refactoring
  - overriding a template
    
- The Library
  - DSLs and Generators
  - Designs
  - Project starter packs

- Hofstadter Developers Corner
  - Implementing Generators
  - Creating DSLs and the accompanying design

- Playbooks and Guides
  - gebbifying an existing project
  - importing types, swagger, and other specs

