# TODO and errors


### Usability

- geb new ...

### Bugs and discords

- collection types
- object vs interface{}
- data-utils/visit 'main' import...?
  - log.go in pkg showing up as `package main`
  - need to specify `module: true` in `pkg: ...`
  - will this get fixed with `gq` and `visit`?
- sub-package logging setup does not pick up third party? ok? manual? how to automate?

### Tests

- InferDataType... in data-utils
- Merge... in data-utils

### Missing or enhancements

- function types
- typedefs
- consts
- enums
- higher-level containers


### Data

- multiple objects in one file
  - yaml (---)
  - json (jsonl)
  - toml (?)
  - xml (?)
- merging dsl and gen files 
- bad design file does not produce an error
- cli's and api's w/o cmds/resources do not generate log/router.go in folders that are still generated
  - this is supposed to be solve with gq boolean expressions and updates to the geb-gen.yaml


### Types

- embed/extend (for RBAC)

