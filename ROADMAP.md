# Order Features Roadmap

### v0.1.0 - Groundwork
----
`[target]: fundamental features`
- basic file structure with tasks only
- execution with only list of commands
- CLI to list and run tasks
- project must include:
    - documentation
    - unit tests
    - examples
    - code quality & development tools/configs

### v0.1.5 - Confidence
----
`[target]: show that the tool is useful, enrich the project with the gained experience`
- use order as its own task runner / build system

### v0.2.0 - Lineage
----
`[target]: inheritance as project signature feature which allows for better task management and separation between definition and implementation`
- inheritance (local):
    - vertical (parent - child):
        - one parent allowed
        - modes: ignore / extend / override
    - horizontal (siblings):
        - with task namespacing/aliasing

### v0.3.0 - Mold
----
`[target]: more complex commands support, parametrization, more real CLI-like experience for tasks execution`
- string templating:
    - task parameters:
        - secrets
        - bool
        - string
- default values
- environmental variables injection

### v0.4.0 - Distance
----
`[target]: definition centralization for multirepo-based projects`
- network-based inheritance:
    - only vertical inheritance
    - authentication support
    - global & task-based checksum system

### v0.5.0 - Singularity
----
`[target]: better management over tree-like structure`
- final form file compilation (to new file):
    - by vertical squashing with all modes applied:
        - ignore: no task present
        - extend: script squash for both tasks
        - override: child only
    - optional `--keep-parent` flag which will allow for **"final file"** to keep parent reference
        - override mode for all tasks as a way to obey the principle o [open-closed](https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle) principle
    - **"final file"** is just normal orderfile.yaml, no special behaviour / syntax

### v0.6.0 - Channel
----
`[target]: pipeline as clear, multi-task scenario definition`
- task pipelining:
    - parameter-based inter-task communication
- artifacts

### v0.7.0 - Vision
----
`[target]: more user friendly tasks/pipelines navigation`
- web interface:
    - local server
    - task & pipelines:
        - visualization
        - edition
        - execution
- introduction of e2e tests for web interface