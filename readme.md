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


## Installation

```
go get github.ibm.com/hofstadter-io/geb
go get github.ibm.com/hofstadter-io/dsl-library
geb sys init
```

(you will need git set up to use `git` rather than `https`)

`git config --global url."git@github.tbm.com:".insteadOf "https://github.ibm.com/"`


## Getting Started

Learn how to [generate a file](./docs/getting-started/file.md)

Build your [first project, a CLI](./docs/getting-started/cli.md)

Get an [overview of Hofstadter](./docs/getting-started/overview.md)

## Documentation

The docs are pretty much m.i.a. but you should still dive in!

Get a feel [here](./docs) for what is to come.
Don't hesitate to ask a question via the GitHub issues either.

[In-depth API walkthrough](./docs/walkthrough) 

[The Concepts](./docs/concepts)

[All about designs](./docs/designs)

[Exploring the DSLs](./docs/explore)

[Developers Documentation](./doc/develop)

[Playbooks and Guides](./docs/guides)

## Projects using Hofstadter

| Project                                                           | types | pkg | api | cli | db  | ci  | notes |
|:--------                                                          |:-----:|:---:|:---:|:---:|:---:|:---:|:------|
| [geb](https://github.ibm.com/hofstadter-io/geb)                   |   x   |  x  |     |  x  |     |     | meta... |
| [Xtalk](https://github.ibm.com/krobots/k8s-csf-xtalk)             |   x   |  x  |     |  x  |     |     | facilitates communication between CSF and Armada |
| [gzi](https://github.ibm.com/hofstadter-io/gzi)                   |   x   |  x  |     |  x  |     |     | GitHub-ZenHub CLI |
| [cego](https://github.ibm.com/hofstadter-io/cego)                 |   x   |  x  |     |  x  |     |     | Tool for visibility at IBM |
| [disgo-frontend-api](https://github.ibm.com/hofstadter-io/geb)    |   x   |  x  |  x  |  x  |  x  |  x  | experimental golang api server |

