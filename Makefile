ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@mkdir -p out
	@cd rust && cargo build --release --quiet
	@go build -o out/bt .

.PHONY: run
run: build
	@go run .

.PHONY: test-rust-lib
test-rust-lib:
	@cd rust && cargo test -- --nocapture

.PHONY: test
test: build
	@go test -v

.PHONY: clean
clean:
	rm -rfv out rust/target
