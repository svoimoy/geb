.Phony: all
all:
	@echo 'Hoftstadter!'

.Phony: build
build:
	@go build

.Phony: clone
clone:
	@cd contracrostipunctus && go build

.Phony: profit
profit:
	@ cp -R contracrostipunctus/cli/golang/* contracrostipunctus/profit
