ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@cd lib/hello && cargo build --release
	@cp lib/hello/target/release/libhello.a lib/
	go build main.go

.PHONY: run
run: build
	@./main

# This is just for running the Rust lib tests natively via cargo
.PHONY: test-rust-lib
test-rust-lib:
	@cd lib/hello && cargo test -- --nocapture

.PHONY: clean
clean:
	rm -rf main_dynamic main_static lib/libhello.a lib/hello/target
