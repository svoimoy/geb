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

```
go get github.ibm.com/hofstadter-io/geb
geb sys init
```

(you will need git set up to use `git` rather than `https`)

`git config --global url."git@github.tbm.com:".insteadOf "https://github.ibm.com/"`


### Usage

The docs are pretty much m.i.a. but you should still dive in!
See the design folders from the projects list below.

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



### Documentation

Coming soon:

- Using geb
    - Getting started
    - Concepts
    - Project Walkthrough
    - The Library
    - Playbooks and Guides
- geb in depth
    - Designs
    - DSLs and Generators
    - Merging and Layering
    - Exploring the DSLs
    - Additional Topics
    - [Hofstadter Developers Corner](./doc/develop)

### Projects using Hofstadter

| Project                                                           | types | pkg | api | cli | notes |
|:--------                                                          |:-----:|:---:|:---:|:---:|:------|
| [geb](https://github.ibm.com/hofstadter-io/geb)                   |   x   |  x  |     |  x  | meta... |
| [Xtalk](https://github.ibm.com/krobots/k8s-csf-xtalk)             |   x   |  x  |  x  |  x  | facilitates communication between CSF and Armada |
| [gzi](https://github.ibm.com/hofstadter-io/gzi)                   |   x   |  x  |     |  x  | GitHub-ZenHug CLI |
| [cego](https://github.ibm.com/hofstadter-io/cego)                 |   x   |  x  |     |  x  | Tool for visibility at IBM |
| [disgo-frontend-api](https://github.ibm.com/hofstadter-io/geb)    |   x   |  x  |  X  |  x  | experimental golang api server |

