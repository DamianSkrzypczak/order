package order

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type RunnerOptions struct {
	NoCommand bool
}

type Runner struct {
	runner *interp.Runner
	opts   RunnerOptions
}

var defaultOptionsForInternalRunner = []interp.RunnerOption{
	interp.Env(expand.ListEnviron(os.Environ()...)),
}

func NewRunner(opts RunnerOptions) (*Runner, error) {
	runner, err := interp.New(defaultOptionsForInternalRunner...)
	if err != nil {
		return nil, err
	}

	return &Runner{runner, opts}, nil
}

func (r *Runner) RunOrder(order *Order) error {
	for _, cmd := range order.Script {
		if err := r.runCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) runCommand(cmd Cmd) error {
	if !r.opts.NoCommand {
		log.Info().Msgf("> %s", cmd)
	}

	parser, err := syntax.NewParser().Parse(strings.NewReader(string(cmd)), "")
	if err != nil {
		return err
	}

	return r.runner.Run(context.Background(), parser)
}
