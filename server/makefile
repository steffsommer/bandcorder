TPARSE_VERSION := 'v0.18.0'

.PHONY: build build-windows livereload test test-nice

build:
	@wails build

build-windows:
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc wails build -platform windows/amd64

livereload:
	wails dev

tests:
	go test ./...

tests-nice:
	@[ "$(tparse -v | awk '{print $3}')" = "$(TPARSE_VERSION)" ] \
		|| go install github.com/mfridman/tparse@$(TPARSE_VERSION)
	@set -o pipefail && go test ./... -json | tparse -all


