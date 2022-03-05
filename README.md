## Introduction

[![GitHub release](https://img.shields.io/github/release/elliotxx/go-cli-prototype.svg)](https://github.com/elliotxx/go-cli-prototype/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/go-cli-prototype/total.svg)](https://github.com/elliotxx/go-cli-prototype/releases)
[![license](https://img.shields.io/github/license/elliotxx/go-cli-prototype.svg)](https://github.com/elliotxx/go-cli-prototype/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/go-cli-prototype.svg)](https://pkg.go.dev/github.com/elliotxx/go-cli-prototype)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/go-cli-prototype/badge.svg)](https://coveralls.io/github/elliotxx/go-cli-prototype)

> This is a cli application with go and cobra.

## Intallation
### Homebrew

The `elliotxx/tap` has macOS and GNU/Linux pre-built binaries available:

```
brew install elliotxx/tap/go-cli-prototype
```

### Build from Source

Starting with Go 1.17, you can install `go-cli-prototype` from source using go install:

```
go install github.com/elliotxx/go-cli-prototype/cmd/go-cli-prototype@latest
```

*NOTE*: This will install `go-cli-prototype` based on the latest available code base. Even though the goal is that the latest commit on the main branch should always be a stable and usable version, this is not the recommended way to install and use `go-cli-prototype`. The version output will show `go-cli-prototype` version (default-version) for go install based builds.


## Usage
Local startup:
```
$ go run cmd/go-cli-prototype/main.go -e hello
hello
$ go run cmd/go-cli-prototype/main.go -V
v0.1.3-9312a46c
```

Local build:
```
$ make build-all
$ ./build/darwin/go-cli-prototype -e hello
hello
$ ./build/darwin/go-cli-prototype -V      
v0.1.3-9312a46c
```

Run all unit tests:
```
make test
```

All targets:
```
$ make help
help                           This help message :)
test                           Run the tests
cover                          Generates coverage report
cover-html                     Generates coverage report and displays it in the browser
format                         Format source code
lint                           Lint, will not fix but sets exit code on error
lint-fix                       Lint, will try to fix errors and modify code
doc                            Start the documentation server with godoc
gen-version                    Update version
clean                          Clean build bundles
build-all                      Build all platforms
build-darwin                   Build for MacOS
build-linux                    Build for Linux
build-windows                  Build for Windows
```