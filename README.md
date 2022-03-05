## Introduction

[![GitHub release](https://img.shields.io/github/release/elliotxx/blueprint.svg)](https://github.com/elliotxx/blueprint/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/blueprint/total.svg)](https://github.com/elliotxx/blueprint/releases)
[![license](https://img.shields.io/github/license/elliotxx/blueprint.svg)](https://github.com/elliotxx/blueprint/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/blueprint.svg)](https://pkg.go.dev/github.com/elliotxx/blueprint)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/blueprint/badge.svg)](https://coveralls.io/github/elliotxx/blueprint)

> This is a cli application with go and cobra.

## Intallation
### Homebrew

The `elliotxx/tap` has macOS and GNU/Linux pre-built binaries available:

```
brew install elliotxx/tap/blueprint
```

### Build from Source

Starting with Go 1.17, you can install `blueprint` from source using go install:

```
go install github.com/elliotxx/blueprint/cmd/blueprint@latest
```

*NOTE*: This will install `blueprint` based on the latest available code base. Even though the goal is that the latest commit on the main branch should always be a stable and usable version, this is not the recommended way to install and use `blueprint`. The version output will show `blueprint` version (default-version) for go install based builds.


## Usage
Local startup:
```
$ go run cmd/blueprint/main.go -e hello
hello
$ go run cmd/blueprint/main.go -V
v0.1.3-9312a46c
```

Local build:
```
$ make build-all
$ ./build/darwin/blueprint -e hello
hello
$ ./build/darwin/blueprint -V      
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