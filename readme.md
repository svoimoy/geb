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


### Quick Introduction

The docs are pretty much m.i.a. but you should still dive in!
See the projects list below for in-depth examples.

#### The project file

Every project begins with a folder containing a `geb.yaml` file and a design directory.

The main fields are:

- name
- about
- output-dir
- design-dir
- dsl-config

The dsl-config sections tells geb where to find DSLs and which generators to use.
It is typical to have all output-dir fields be "." unless you are generating
the same dsl in multiple languages (experimental).

The `common`, `type`, and `pkg` DSLs are almost always included.
(until the dependent generators feature is added)

```
name: "geb"
about: "The geb command-line tool for making profit."

output-dir: "."

design-dir: "design"

dsl-config:
  paths:
    - "dsl"
    - "$GOPATH/src/github.ibm.com/hofstadter-io/geb/dsl"
  default:
    - dsl: common
      gen:
        - golang
      output-dir: "."
    - dsl: type
      gen:
        - golang
      output-dir: "."
    - dsl: pkg
      gen:
        - golang
      output-dir: "."
    - dsl: cli
      gen:
        - golang
      output-dir: "."

log-config:
  default:
    level: warn
    stack: false
```

#### The design layout

The typical layout for a project is defined by the design layout.

CLI:

```
my-project/
    geb.yaml
    design/
        
        cli.yaml

        commands/
            <command-name>.yaml
            ...
            <command-name>/             (when it has subcommands)
                <sub-command-name>.yaml
                ...
            
            
        lib/
            app-specific/
                pkg.yaml
                subdir/pkg.yaml
                ...
            types/
                type.yaml
                subdir/type.yaml
                ...
```

#### Generated output layout

When you run `geb gen`,
the output structure will align with the design structure.
You can use any folder layout you wish to organize tyour code.
Note, the APIs and CLIs will generate folder structure which matches their definition.

```

my-project/

    main.go

    commands/
        root.go
        <command-name>.go
        ...

        <command-name>/             (when it has subcommands)
            <sub-command-name>.yaml
            ...
            
            <sub-command-name>/             (when it has sub-sub-commands)
                <sub-sub-command-name>.yaml
                ...

        lib/

            app-specific/
                <name>PubPkg.yaml
                ...
                subdir/<name>PubPkg.yaml
                ...

            types/
                type<Name>.yaml
                subdir/type<Name>.yaml
                ...
```

Types and Packages will land in a folder which aligns with the design folder path.


#### A first design

Let's make a hello world CLI.
Create a folder with the name `hello`,
add the `geb.yaml`, and a `design/cli.yaml`.

You can copy the `geb.yaml` file from above, making sure to change the name to `hello`

Now put the following in the `cli.yaml`:

```
cli:
  name: hello
  short: "A simple hello world cli."
```

#### A first generation

Now run:

```
# generate output from design
geb gen

# make the code pretty (it's in golang)
gofmt -w .

# build the hello cli
go build

# run the command
./hello
```

pretty boring eh?


#### Adding to the code

Now let's change the existing command
and have it do something.

Add pkg `"fmt"` in the imports:

```
import (
	// HOFSTADTER_START import
    "fmt"
	// HOFSTADTER_END   import
    ...
```

Add print the arguments to the command like so:

```
// HOFSTADTER_START cmd_run
fmt.Println("Hello! ", args)
// HOFSTADTER_END   cmd_run
```



#### Regenerating and rebuilding

Run the same commands as before again:

```
# generate output from design
geb gen

# make the code pretty (it's in golang)
gofmt -w .

# build the hello cli
go build

# run the command
./hello douglas
./hello
./hello 1arg 2arg ah ah ah
```

When you change your design and (re)generate the output,
geb takes care not to distrub your code.
All you have to do is edit between the `HOFSTADTER_*` tags.


#### Adding a command

Add a command to the `design/cli.yaml` file:

```
cli:
  name: hello
  short: "A simple hello world cli."

  commands:
    - name: there
      short: "say something to someone"
      args:
        - name: who
          type: string
          required: true
        - name: what
          type: string
```

Now regenerate, rebuild, and then run:
`hello there <name> <message>`.

Try updating the print message to be
formatted nicer and say something
when no message is supplied.


### Documentation

Coming soon...

Get a feel [here](./docs) for what is to come.
Don't hesitate to ask a question via the GitHub issues either.

An [API walkthrough](./docs/intro/api.md) has been started.

### Projects using Hofstadter

| Project                                                           | types | pkg | api | cli | notes |
|:--------                                                          |:-----:|:---:|:---:|:---:|:------|
| [geb](https://github.ibm.com/hofstadter-io/geb)                   |   x   |  x  |     |  x  | meta... |
| [Xtalk](https://github.ibm.com/krobots/k8s-csf-xtalk)             |   x   |  x  |  x  |  x  | facilitates communication between CSF and Armada |
| [gzi](https://github.ibm.com/hofstadter-io/gzi)                   |   x   |  x  |     |  x  | GitHub-ZenHub CLI |
| [cego](https://github.ibm.com/hofstadter-io/cego)                 |   x   |  x  |     |  x  | Tool for visibility at IBM |
| [disgo-frontend-api](https://github.ibm.com/hofstadter-io/geb)    |   x   |  x  |  X  |  x  | experimental golang api server |

