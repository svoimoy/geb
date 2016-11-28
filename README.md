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

