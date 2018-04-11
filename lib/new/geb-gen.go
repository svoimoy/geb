package new

var GEN_DIRS = []string{"designs", "partials", "templates", "new", "static"}

const GEN_TEMPLATE = `name: "{{name}}"
about: "All about {{#if about}}{{about}}{{else}}{{name}}{{/if}}..."
version: "0.0.1"
language: golang
type: "{{type}}"

# Below is example content, update to your needs.
# See more at [doc link t.b.d.]

dependencies:
  generators:
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

template-configs:

  - name: once-files
    field: "."
    templates:
      - in: "main.go"
        out: "main.go"
      - in: "root.go"
        out: "commands/root.go"
      - unless: commands
        in: "rootonlylog.go"
        out: "commands/log.go"
      - when: commands
        in: "log.go"
        out: "commands/log.go"

  - name: command
    field: commands
    templates:
      - in: "cmd.go"
        out: "commands/{{{string "{{trimto_first pkgPath '/' false}}"}}}.go"
      - when: commands
        in: "log.go"
        out: "commands/{{string "{{name}}"}}/log.go"
`
