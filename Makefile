# get current dir, set design and out paths
design_path="$(shell pwd | sed "s|^${GOPATH}/src/||")/design"
output_path="."

.PHONY: default
default: clean hofstadter

.PHONY: clean
clean:
	@rm -rf $(output_path)/{dsl,impl,gen}

.PHONY: hofstadter
hofstadter:
	@echo "output to: '$(output_path)'"
	@goagen --out $(output_path) -d $(design_path) gen --pkg-path=github.com/hofstadter-io/hofstadter/gen/dsl

