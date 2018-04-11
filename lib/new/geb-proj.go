package new

var PACKAGE_DIRS = []string{"design", "dsl", ".geb/shadow", ".geb/tmp"}

const PACKAGE_TEMPLATE = `name: {{name}}
about: "All about {{#if about}}{{about}}{{else}}{{name}}{{/if}}..."

output-dir: "."

design-dir: "design"

# Below is example content, update to your needs.
# See more at [doc link t.b.d.]

dsl-config:
  paths:
    - "./dsl"
  default:
    - dsl: cli
      gen:
        - golang
      output-dir: "."

run-config:
  all:
    - name: "all"
      command: "geb run"
      args:
        - gen
        - fmt
        - install
  regen:
    - name: "regen"
      command: "geb run"
      args:
        - gen
        - fmt
  gen:
    - name: generate
      command: "geb"
      args:
        - "gen"
  fmt:
    - name: format
      command: "gofmt -w main.go commands engine lib"
  build:
    - name: build
      command: "go build"
  install:
    - name: install
      command: "go install"


log-config:
  default:
    level: warn
    stack: false
`
