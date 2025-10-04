# README

## About

This project is the server/desktop application of Bandcorder, which can be controlled via
the app additionally to its own UI. The project utilizes the cross-platform application
framework [wails](https://wails.io/).

The backend is written in Go. It includes a REST API, which takes commands from the
app clients. The frontend is a React web application.

## Prerequisites

- `go version` >= 1.25
- `wails version` >= 2.10.2
- `node -v` >= 20.11.0
- `make` (available on all major platforms)
- A C compiler, e.g. `gcc` (see platform-specific requirements below)

### C Compiler Requirements

**This project requires CGO** because it uses low-level audio libraries. CGO needs a C compiler to build the native code portions.

#### On Windows

Windows does not come with a C compiler. You must install one of the following:

- **MinGW-w64** (recommended): Download from [mingw-w64.org](https://www.mingw-w64.org/) or install via package managers:

  ```powershell
  # Via Chocolatey
  choco install mingw

  # Via MSYS2
  pacman -S mingw-w64-x86_64-gcc
  ```

After installation, ensure `gcc.exe` or `clang.exe` is in your `PATH`. CGO will auto-detect it.

#### On Linux (Native Builds)

Most Linux distributions include `gcc` by default. If not:

```bash
# Debian/Ubuntu
sudo apt install build-essential

# Arch Linux
sudo pacman -S base-devel
```

#### On Linux (Cross-Compiling to Windows)

To build Windows executables from Linux, install MinGW-w64:

```bash
# Arch Linux
sudo pacman -S mingw-w64-gcc

# Debian/Ubuntu
sudo apt install gcc-mingw-w64
```

## Make Targets

| Target                     | Description                                                                                            |
| -------------------------- | ------------------------------------------------------------------------------------------------------ |
| `make livereload`          | Compiles the project and hot-reloads on changes. To debug the frontend, go to `http://localhost:34115` |
| `make build`               | Compile the application for your current OS and architecture                                           |
| `make cross-build-windows` | Cross-compile the application for 64-bit Windows using MinGW (Linux only)                              |
| `make test`                | Run all Go tests in the project                                                                        |
