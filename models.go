package order

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

// Orderfile represents single instance of orderfile
type Orderfile struct {
	Version string `yaml:"version"`
	Orders  Orders `yaml:"orders"`
}

func (o *Orderfile) ListOrdersNames() []string {
	orders := []string{}
	for orderName := range o.Orders {
		orders = append(orders, orderName)
	}

	return orders
}

// Orders represents tasks section of file
type Orders map[string]Order

// Order represents single task - unit of order execution
type Order struct {
	Description string `yaml:"description"`
	Script      []Cmd  `yaml:"script"`
}

func (o *Order) Run(hideCommands bool) error {
	r, err := interp.New(
		interp.Dir("."),
		interp.Env(expand.ListEnviron(os.Environ()...)),
		interp.OpenHandler(func(ctx context.Context, path string, flag int, perm os.FileMode) (io.ReadWriteCloser, error) {
			return interp.DefaultOpenHandler()(ctx, path, flag, perm)
		}),
		interp.StdIO(os.Stdin, os.Stdout, os.Stderr),
	)
	if err != nil {
		return err
	}

	for _, cmd := range o.Script {
		if err := cmd.Run(r, hideCommands); err != nil {
			return err
		}
	}

	return nil
}

// Cmd represents single script command
type Cmd string

func (cmd Cmd) Run(r *interp.Runner, hideCommand bool) error {
	if !hideCommand {
		log.Info().Msgf("> %s", cmd)
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(string(cmd)), "")
	if err != nil {
		return err
	}

	return r.Run(context.Background(), p)
}
