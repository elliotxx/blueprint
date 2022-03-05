## 简介

[![GitHub release](https://img.shields.io/github/release/elliotxx/blueprint.svg)](https://github.com/elliotxx/blueprint/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/blueprint/total.svg)](https://github.com/elliotxx/blueprint/releases)
[![license](https://img.shields.io/github/license/elliotxx/blueprint.svg)](https://github.com/elliotxx/blueprint/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/blueprint.svg)](https://pkg.go.dev/github.com/elliotxx/blueprint)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/blueprint/badge.svg)](https://coveralls.io/github/elliotxx/blueprint)

> 这是一个用 Go 和 cobra 编写的 CLI 应用

## 📜 语言

[English](https://github.com/elliotxx/blueprint/blob/master/README.md) | [简体中文](https://github.com/elliotxx/blueprint/blob/master/README-zh.md)

## 🛠️ 安装

### Homebrew

`elliotxx/tap` 有 MacOS 和 GNU/Linux 的预编译二进制版本可用:

```
brew install elliotxx/tap/blueprint
```

### 从源码构建

使用 Go 1.17+ 版本，你可以通过 `go install` 直接从源码安装 `blueprint`:

```
go install github.com/elliotxx/blueprint/cmd/blueprint@latest
```

*注意*: 你将基于代码仓库最新的可用版本安装 `blueprint`，尽管主分支的最新提交应该始终是一个稳定和可用的版本，但这不是安装和使用 `blueprint` 的推荐方式。通过 `go install` 安装的 `blueprint` 版本输出将显示默认版本号（default-version）。

## ⚡ 使用

本地启动：

```
$ go run cmd/blueprint/main.go -e hello
hello
$ go run cmd/blueprint/main.go -V
v0.1.3-9312a46c
```

本地构建：

```
$ make build-all
$ ./build/darwin/blueprint -e hello
hello
$ ./build/darwin/blueprint -V      
v0.1.3-9312a46c
```

运行所有单元测试：

```
make test
```

全部可用的 Target:

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
