+++
title = "Installation"
weight = 2
chapter = false
+++

{{% notice info %}}
Currently Order supports only installation via `go get`
but binary releases and support for various package managers is comming soon.
{{% /notice %}}

#### Using `go get` (installed go required)
To install "order" tool run:
```
go get -u github.com/DamianSkrzypczak/order
```
then check if everything works file with
```
order --help
```
you should see usage similar to that:
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
