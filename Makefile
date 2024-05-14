ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@mkdir -p out
	@cd rust && cargo build --release
	@go build -o out/btscan .

.PHONY: run
run: build
	@./out/btscan

# This is just for running the Rust lib tests natively via cargo
.PHONY: test-rust-lib
test-rust-lib:
	@cd rust && cargo test -- --nocapture

.PHONY: clean
clean:
	rm -rfv out rust/target
