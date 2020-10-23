# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2020-10-24

### Added
- Features:
    - Orderfile v0.1.0 features (see [ROADMAP.md v0.1.0 section](ROADMAP.md#v010---groundwork))
- Documentation:
    - [Project documentation](https://damianskrzypczak.github.io/order)
    - [README.md](README.md) with project description / instalation guide etc.
    - Contribution guide [CONTRIBUTING.md](CONTRIBUTING.md)
    - This changelog
- Quality assurance:
    - [Github actions configuration](https://github.com/DamianSkrzypczak/shift/actions)
    - More golangci-lint restrictions (full code documentation requirement)
    - Tests (with 100% coverage)
- Other:
    - [Orderfile.yml](Orderfile.yml) (order now uses itself to run project tasks)

### Changed
- Project structure:
    - now main package is kept in project root
        - this allows for command installation with `go get -u github.com/DamianSkrzypczak/order`
    - order components code is defined within [internal](internal) directory

### Removed
- Examples directory

## [0.0.0] - 2020-10-18
### Added
- Basic project structure with:
    - Go package structure
    - Developer tool configs like:
        - [.editorconfig](.editorconfig)
        - [.golangci.yml](.golangci.yml)
    - [ROADMAP.md](ROADMAP.md#v010---groundwork) with information about planned features
    - [LICENSE](LICENSE) file