+++
title = "Roadmap"
weight = 3
chapter = false
disableNextPrev = true
+++

### v0.1.0 - Groundwork ⛏️
----
🎯 **objective**: fundamental features
- basic file structure with orders only
- execution with only list of commands
- CLI to list and run orders
- project must include:
    - documentation
    - unit tests
    - code quality & development tools/configs
- use order as its own order runner / build system

### v0.2.0 - Gears ⚙️
----
🎯 **objective**: more complex commands support, parametrization, more real CLI-like experience for orders execution
- add basic Orderfile.yml initialization with --init flag
- referencing one order in another
- string templating:
    - order parameters/flags:
        - secrets
        - bool
        - string
    - default values
    - "passthrough" option:
        - will pass all flags/positional agrs to marked place in order cmd string
        - passtrhough allows for special "help-cmd" which will be helpful to see command's original help
- environmental variables injection

### v0.3.0 - Lineage 👪
----
🎯 **objective**: inheritance as project signature feature which allows for better order management and separation between definition and implementation`
- inheritance (local):
    - vertical (parent - child):
        - one parent allowed
        - modes: ignore / extend / override
    - horizontal (siblings):
        - with order namespacing/aliasing

### v0.4.0 - Distance 📏
----
🎯 **objective**: definition centralization for multirepo-based projects
- network-based inheritance:
    - only vertical inheritance
    - authentication support
    - global & order-based checksum system

### v0.5.0 - Singularity 🔮
----
🎯 **objective**: better management over tree-like structure
- final form file compilation (to new file):
    - by vertical squashing with all modes applied:
        - ignore: no order present
        - extend: script squash for both orders
        - override: child only
    - optional `--keep-parent` flag which will allow for **"final file"** to keep parent reference
        - override mode for all orders as a way to obey the principle o [open-closed](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle) principle
    - **"final file"** is just normal Orderfile.yml, no special behaviour / syntax

### v0.6.0 - Channel 🚿
----
🎯 **objective**: pipeline as clear, multi-order scenario definition
- order pipelining:
    - parameter-based inter-order communication
- artifacts

### v0.7.0 - Vision 🔭
----
🎯 **objective**: more user friendly order/pipelines navigation
- web interface:
    - local server
    - order & pipelines:
        - visualization
        - edition
        - execution
- introduction of e2e tests for web interface