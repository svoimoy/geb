# get current dir, set design and out paths
design_path="$(shell pwd | sed "s|^${GOPATH}/src/||")/design"
output_path="."

.PHONY: default
default: clean hofstadter

.PHONY: clean
clean:
	@rm -rf $(output_path)/cli

.PHONY: hofstadter
hofstadter:
	@echo "output to: '$(output_path)'"
	@goagen --out $(output_path) -d $(design_path) gen --pkg-path=github.com/hofstadter-io/geb-dsl-cli/gen/dsl

.PHONY: cli
cli:
	go build -o geb cli/*.go
