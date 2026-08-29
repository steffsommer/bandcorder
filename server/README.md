# README

## About

This project is the server/desktop application of Bandcorder, which can be controlled via
the app additionally to its own UI. The project utilizes the cross-platform application
framework [wails](https://wails.io/).

The backend is written in Go. It includes a REST API, which takes commands from the
app clients. The frontend is a React web application.

## Prerequisites

The development environment is managed with [devbox](https://www.jetify.com/devbox), which
pins all system-level dependencies (Go, `pkg-config`, GTK/WebKit, Mesa, and the mingw
cross-compiler used for Windows builds). See [`devbox.json`](../devbox.json) at the
repository root for the exact package set.

To get a fully provisioned shell with everything on your `PATH`:

```bash
# Install devbox (once): https://www.jetify.com/docs/devbox/installing_devbox/

# From the repository root, enter the environment:
devbox shell
```

Entering the devbox shell also installs the pinned `wails` CLI automatically via the
`init_hook` in [`devbox.json`](../devbox.json) (`go install ...@$WAILS_VERSION`). To bump
wails, change `WAILS_VERSION` in `devbox.json` and keep it in sync with
`github.com/wailsapp/wails/v2` in `go.mod`. `node` (for the wails frontend build) is expected
to be available on your system; `make tests-nice` and `make gen-mocks` install their extra Go
tools (`tparse`, `mockery`) on demand.

## Make Targets

| Target                     | Description                                                                                                |
| -------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `make livereload`          | Compiles the project and hot-reloads on changes. To debug the frontend, go to `http://localhost:34115`     |
| `make build`               | Compile the application for your current OS and architecture                                               |
| `make cross-build-windows` | Cross-compile the application for 64-bit Windows using MinGW (Linux only)                                  |
| `make tests`               | Run all Go tests                                                                                           |
| `make tests-nice`          | Run all Go tests with nicer output, requires binaries installed via `go install` to be available in `PATH` |
