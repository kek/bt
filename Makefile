ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@cd rust && cargo build --release
	@cp rust/target/release/libhello.a lib/
	go build main.go

.PHONY: run
run: build
	@./main

# This is just for running the Rust lib tests natively via cargo
.PHONY: test-rust-lib
test-rust-lib:
	@cd rust && cargo test -- --nocapture

.PHONY: clean
clean:
	rm -rf main_dynamic main_static lib/libhello.a rust/target
