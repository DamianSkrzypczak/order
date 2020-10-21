package runner

import (
	"context"
	"os"
	"strings"

	"github.com/DamianSkrzypczak/order/internal/orderfile"
	"github.com/rs/zerolog/log"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

// Options represents set of all available runner options
type Options struct {
	NoCommand bool
}

// Runner represents order command runner
type Runner struct {
	runner *interp.Runner
	opts   Options
}

var defaultOptionsForInternalRunner = []interp.RunnerOption{
	interp.Env(expand.ListEnviron(os.Environ()...)),
	interp.StdIO(os.Stdin, os.Stdout, os.Stderr),
}

// NewRunner produces Runner instance configured with given options
func NewRunner(opts Options) (*Runner, error) {
	runner, err := interp.New(defaultOptionsForInternalRunner...)
	if err != nil {
		return nil, err
	}

	return &Runner{runner, opts}, nil
}

// RunOrder executes all commands defined in given order
// within single environment shared between all commands
func (r *Runner) RunOrder(order *orderfile.Order) error {
	for _, cmd := range order.Script {
		if err := r.runCommand(cmd); err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) runCommand(cmd orderfile.Cmd) error {
	if !r.opts.NoCommand {
		log.Info().Msgf("> %s", cmd)
	}

	parser, err := syntax.NewParser().Parse(strings.NewReader(string(cmd)), "")
	if err != nil {
		return err
	}

	return r.runner.Run(context.Background(), parser)
}
