# Developing GEB

This document is intended for those interested
in contributing to the geb source code.
Use the following links if you are interested in:

- [Developing DSLs and Generators]
- [Developing applications with GEB]


### Setup

To develop geb, fork and clone the repositories

```
git clone git@github.com:hofstadter-io/geb
git clone git@github.com:hofstadter-io/dsl-library

cd geb
go build
```

and then enter development mode with

```
geb sys dev
```

This will add some symlinks so the development
geb and dsl-library are available system wide.
You can then test your changes against
a project in any directory.


### Adding new dependencies

GEB's dependencies are vendored and committed to the repository.
This is done to prevent automated build failures when GitHub
or another code host is down.

GEB uses [dep](https://github.com/golang/dep) to manage dependencies.
Add a dependency with `dep ensure -add <location>` where location
is like a `go get ...`.

#### Take Note:

Everytime `dep ensure ...` is run, it wants to update the yaml.v2 library.
The `gopkg.in/yaml.v2` has a slight modification
to make the default map type a `map[string]interface{}` rather
than having the key be and interface as well.
You will almost certainly need to run
`git checkout vendor/gopkg.in/yaml.v2/decode.go`
after using dep.

