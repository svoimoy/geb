# geb - The Hofstadter tool

__geb__ is the command-line tool for hofstadter
data-centric design and code generation.
With it, you can:

- Create designs and generate code
- Modify code and regenerate without disruption
- Customize templates and extend generators for your particular use-case
- Share designs and templates across your applications and organization.
- Contribute to the ecosystem by sharing your designs, templates, generators, or complete applications.

### Installation

Download a binary from the [releases](https://github.com/hofstadter-io/geb/releases).

or get the latest by `go get`'n from github:

`go get github.com/hofstadter-io/geb`

Then do the initial setup:
`geb initial-setup`.
This will download the
default designs and templates
from the geb-hub library.


### Documentation


### Developing

Please contribute, we are working towards and ecosystem.

##### Tips:

Make the following symlinks,
from the development folder:

1. the `geb` binary to `/usr/local/bin/geb`
1. the `dsl` folder go `~/.hofstadter/dsl`

This will keep you defaulting to
the latest build and templates,
when working in any directory
and letting you override
locally while still developing.

