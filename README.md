![status](https://img.shields.io/badge/status-beta-yellow.svg)
[![go report](https://goreportcard.com/badge/github.com/DamianSkrzypczak/order)](https://goreportcard.com/report/github.com/DamianSkrzypczak/order)
[![actions status](https://github.com/DamianSkrzypczak/order/workflows/Testing/badge.svg)](https://github.com/DamianSkrzypczak/order/actions)
[![documentation](https://img.shields.io/badge/documentation-reference-%234DB6AC)](https://damianskrzypczak.github.io/order)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c)](https://pkg.go.dev/github.com/DamianSkrzypczak/order)
[![godoc](https://godoc.org/github.com/DamianSkrzypczak/order?status.svg)](http://godoc.org/github.com/DamianSkrzypczak/order)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/DamianSkrzypczak/order/blob/master/LICENSE)


<!-- <img align="right" height="100px" src="https://raw.githubusercontent.com/DamianSkrzypczak/order/master/media/logo.png"> -->
<!-- <img alt="logo" align="right" width="350px" src="./media/logo.png"> -->

# Order
✔ A modern approach to running your tasks

> :warning: **Please be warned that the project is early in its development, which means the API and code base are generally unstable for the short term.**

## Introduction

Order is a versatile task runner / build system created to assist with the project development and maintance processes.

It's main goal is to give to the developer:
- script definition file with clear and intuitive structure
- task-based subcommands
    - **[not yet implemented]** with support for user-defined flagset
- **[not yet implemented]** task inheritance which provides:
    - centralization of common task code (multirepo project support)
    - ability to divide tasks into file-based, aliased namespaces
    - support for parent task checksum validation
    - optional separation between definition and implementation
    - compliation of whole inheritance tree to single file

See [ROADMAP](https://damianskrzypczak.github.io/order/features/roadmap/) for more details on feature development progress.

## Installation
### Latest version from Source
```bash
go get -u github.com/DamianSkrzypczak/order
```

## Usage
```
Usage: order [options...] <order-name>

Options:
      --debug         debug mode
  -l, --list          list orders
      --no-color      do not color the output
      --no-command    hide currently executed command
      --no-level      hide logging level
  -p, --path string   path to orderfile (default "./Orderfile.yml")
      --version       print version of orderfile (and if loaded, Orderfile.yml)
```

## Documentation
See [damianskrzypczak.github.io/order](https://damianskrzypczak.github.io/order/) for the documentation.

## Contribution & Code of conduct
See [CONTRIBUTING.md](CONTRIBUTING.md)
and our [CODE_OF_CONDUCT](CODDE_OF_CONDUCT.md) for more details on contribution process.

### Conventions
This project utilizes:
- [Semantic Versioning](https://semver.org/spec/v2.0.0.html) for versioning
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages
- [Keep a changelog](https://keepachangelog.com/en/1.0.0/) for CHANGELOG
- [Contributor Covenant](https://www.contributor-covenant.org) for code of conduct

## Credit & License
Order is licensed under the terms of the MIT license. You can find the complete text in [LICENSE](LICENSE).

Please refer to the Git commit log for a complete list of contributors.