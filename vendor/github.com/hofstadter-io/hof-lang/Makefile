.Phony: all
all: parser build
	@echo 'Hoftstadter!'

.Phony: parser
parser: lib/parser/hof.go
	@echo 'building lang'

lib/parser/hof.go: lib/parser/hof.peg
	@pigeon -o lib/parser/hof.go lib/parser/hof.peg

.Phony: build
build:
	@go build

.Phony: gen
gen:
	@geb gen
	@gofmt -w main.go commands

.Phony: geb
geb: gen
	@echo "geb!"

.Phony: fmt
fmt:
	@gofmt -w main.go commands
	
.Phony: cloc
cloc:
	@cloc design
	@cloc subdesigns
	@cloc main.go commands

