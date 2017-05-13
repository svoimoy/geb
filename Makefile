.Phony: all
all:
	@echo 'Hoftstadter!'

.Phony: build
build:
	@go build

.Phony: gen
gen:
	@geb gen
	@gofmt -w main.go commands engine lib

.Phony: geb
geb: gen
	@echo "geb!"

.Phony: fmt
fmt:
	@gofmt -w main.go commands engine lib
	
.Phony: clone
clone:
	@cd contracrostipunctus && go build

.Phont: poor
poor:
	@rm -rf profit

.Phony: profit
profit: poor
	@find contracrostipunctus -type f \! -name "contracrostipunctus" -exec sed -i '' "s/geb\/contracrostipunctus/geb/g" {} \;
	@mkdir -p profit
	@cp -R contracrostipunctus/* profit
	@geb gen > /dev/null

.Phony: serious-profit
serious-profit:
	@rm profit/contracrostipunctus
	@cp -fR profit/* .

.Phony: cloc
cloc:
	@cloc design
	@cloc main.go commands engine lib

