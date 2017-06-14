.Phony: all
all:
	@echo 'Hoftstadter!'

.Phony: build
build:
	@go build

.Phony: gen
gen:
	@geb gen
	@gofmt -w io manip visit

.Phony: geb
geb: gen
	@echo "geb!"

.Phony: fmt
fmt:
	@gofmt -w io manip visit
	
.Phony: cloc
cloc:
	@cloc design
	@cloc io manip visit

