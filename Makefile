ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@mkdir -p src/build
	@cd src/build && cmake .. && make
	@go build -o build/bt .

.PHONY: run
run: build
	@go run .

.PHONY: test
test: build
	@go test -v

.PHONY: clean
clean:
	rm -rfv rust/target build
