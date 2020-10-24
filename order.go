package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"

	"github.com/DamianSkrzypczak/order/internal/orderfile"
	"github.com/DamianSkrzypczak/order/internal/runner"
	"github.com/DamianSkrzypczak/order/preorders/completion"
)

func printOrders(o *orderfile.Orderfile) {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 8, 1, '\t', tabwriter.AlignRight)

	for _, name := range o.ListOrdersNames() {
		ord, ok := o.GetOrder(name)
		if ok {
			lines := strings.Split(ord.Description, "\n")
			fmt.Fprintf(writer, "%s\t%s\n", name, lines[0])

			if len(lines) > 1 {
				for _, line := range lines[1:] {
					fmt.Fprintf(writer, "%s\t%s\n", "", line)
				}
			}
		}
	}

	err := writer.Flush()
	if err != nil {
		log.Error().Msgf("Could not list orders due to error: %s", err)
		return
	}

	log.Info().Msgf("Available orders:\n%s", strings.TrimSuffix(buf.String(), "\n"))
}

const help = `Usage: order [options...] <order-name>

Options:
`

func main() {
	orderCmd := pflag.NewFlagSet("order", pflag.ContinueOnError)
	orderCmd.SetInterspersed(false)

	// output adjustements
	noCommand := orderCmd.Bool("no-command", false, "hide currently executed command")
	noLogLevel := orderCmd.Bool("no-level", false, "hide logging level")
	noColor := orderCmd.Bool("no-color", false, "do not color the output")

	// core flags
	debugModeOn := orderCmd.Bool("debug", false, "debug mode")
	listOrders := orderCmd.BoolP("list", "l", false, "list orders")
	orderfilePath := orderCmd.StringP("path", "p", "./Orderfile.yml", "path to orderfile")
	printVersion := orderCmd.Bool("version", false, "print version of orderfile (and if loaded, Orderfile.yml)")

	// (hiden) completion adding flags
	addBashCompletion := orderCmd.Bool("add-bash-completion", false, "add bash completion")
	_ = orderCmd.MarkHidden("add-bash-completion")

	orderCmd.Usage = func() {
		log.Logger = newHelpLogger()
		log.Info().Msgf("%s%s", help, strings.Trim(orderCmd.FlagUsages(), "\n"))
	}

	err := orderCmd.Parse(os.Args[1:])

	log.Logger = newLogger(*debugModeOn, *noLogLevel, *noColor)

	if err == pflag.ErrHelp {
		return
	}

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	ordfile, err := orderfile.NewOrderFileFrom(*orderfilePath)

	if *printVersion {
		log.Info().Msgf("Order version: %s", Version)

		if err == nil {
			log.Info().Msgf("%s version: %s", *orderfilePath, ordfile.Version)
		}

		return
	}

	if err != nil {
		log.Error().Msgf("Could not load orderfile from %s\n due to error: %s", *orderfilePath, err)
		os.Exit(1)
	}

	if *listOrders {
		printOrders(ordfile)
		return
	}

	var selectedOrder *orderfile.Order

	var orderName string

	if *addBashCompletion {
		orderName = `builtin order "AddBashCompletion"`
		selectedOrder = completion.BashCompletionOrder
	} else {
		orderName = orderCmd.Arg(0)
		if orderName == "" {
			log.Error().Msg("No order specified, exiting")
			os.Exit(1)
		}

		ord, ok := ordfile.GetOrder(orderName)
		selectedOrder = ord
		if !ok {
			log.Error().Msgf(`Could not find order "%s"`, orderName)
			printOrders(ordfile)
			os.Exit(1)
		}
	}

	r, err := runner.NewRunner(
		runner.Options{
			NoCommand: *noCommand,
		},
	)

	if err != nil {
		log.Error().Msgf("Failed to create runner due to %s", err)
		os.Exit(1)
	}

	if err := r.RunOrder(selectedOrder); err != nil {
		log.Error().Msgf("Execution of order \"%s\" failed with: \n%s", orderName, err)
		os.Exit(1)
	}
}
