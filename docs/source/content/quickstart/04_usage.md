+++
title = "Usage"
weight = 4
chapter = false
+++

{{% notice info %}}
Currently Order doesn't support flags/parameters for subcommand.\
However, this feature is planned in our [ROADMAP]({{< ref "/roadmap/#v020---gears-" >}} "ROADMAP") will be added very soon.
{{% /notice %}}

#### Checking version
Running `order --version` will display:
- current version of Order tool
- Orderfile.yml version (if present in current directory or pointet with `--path` flag)
```bash
# With Orderfile.yml present:
> order --version
INF Order version: 0.1.0
INF ../Orderfile.yml version: 0.1.0

# Without Orderfile.yml present or pointed with --path:
> order --version
INF Order version: 0.1.0
```


#### Getting help
Running `order -h` will display help message with all available flags
```bash
> order -h
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

#### Listing available orders
Running `order -l` will list all orders that are available in context of loaded Orderfile.yml
```bash
# Example for Orderfile.yml used in order project
> order -l
INF Available orders:
add-autocompletion
build                   build order binary into ./bin/ directory
build-docs              build project documentation
precommit               run all necessary precommit checks:
                          - goimports (fix-in-place)
                          - gofmt (fix-in-place)
                          - golangci-lint (with .golangci.yml config)
                          - go mod tidy & verify
                          - run unit tests
rm-autocompletion
serve-docs              serve project documentation locally
setup-dev               setup development tools
setup-doc               setup documentation tools:
                          - hugo (into ./bin/ directory)
                          - hugo-theme-learn (git submodule)
test                    run go test with coverage report
test-html               run go test with HTML coverage report
```

#### Using Orderfile from different location
By default, order expects `Orderfile.yml` to be present in current working directory,
however we can point to other file by running
```
order -p=<path>
```

