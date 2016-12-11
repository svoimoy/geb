.Phony: all
all:
	@echo 'Hoftstadter!'

.Phony: build
build:
	@go build

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
