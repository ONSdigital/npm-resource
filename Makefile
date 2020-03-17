BUILD=build

.PHONY: all
all: build package

.PHONY: build
build:
	go build -o $(BUILD)/built-out ./out/cmd/out
	go build -o $(BUILD)/built-check ./check/cmd/check
	go build -o $(BUILD)/built-in ./in/cmd/in

.PHONY: package
package:
	docker build .
