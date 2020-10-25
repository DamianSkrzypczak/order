+++
title = "Adding command line completion"
weight = 2
chapter = false
+++

### Supported shells
Order offers supports command line autocompletion for selected shells:
- bash

### Adding completion
Order completion is done by using **hidden flags**.

Each flag execute special order which defines all the necessary steps to ensure completion for chosen shell.

{{% notice note %}}
For transparency, pre-defined orders behaves in same fashion that user-defined orders would do, this means that for default CLI parameters, all commands will be printed to stdout during task execution.
{{% /notice %}}


#### BASH
To add bash completion for order commands run:
```
order --add-bash-completion
```
When `--no-comamand` flag present, output should include all executed commands.


{{% notice info %}}
Remember to run `source ~/.bashrc` or create new terminal session
to load autocompletion after sucessful registration
{{% /notice %}}


