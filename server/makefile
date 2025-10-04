.PHONY: test build-windows

livereload:
	wails dev

test:
	go test ./...

build:
	wails build

build-windows:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc wails build -platform windows/amd64
