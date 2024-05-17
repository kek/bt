ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build run

.PHONY: build
build:
	@mkdir -p out
	@cd rust && cargo build --release
	@mkdir -p build
	@cp rust/target/release/libbluetooth.a build
	@cp `find rust/target/release/build -name '*.a'` build
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
	rm -rfv out rust/target build
