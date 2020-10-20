package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"

	"github.com/DamianSkrzypczak/order"
)

func setupLogger(debug bool) {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}
	log.Logger = log.Output(output)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug mode: on")
	}
}

func main() {
	orderCmd := pflag.NewFlagSet("order", pflag.ContinueOnError)
	orderCmd.SetInterspersed(false)

	debugModeOn := orderCmd.BoolP("debug", "d", false, "debug mode")
	listOrders := orderCmd.BoolP("list", "l", false, "list orders")
	noCommands := orderCmd.BoolP("no-commands", "n", false, "do not show currently executed command")
	orderfilePath := orderCmd.StringP("path", "p", "Orderfile.yml", "path to orderfile")

	err := orderCmd.Parse(os.Args[1:])

	if err != nil {
		log.Error().Err(err).Send()
		os.Exit(1)
	}

	if err == pflag.ErrHelp {
		return
	}

	setupLogger(*debugModeOn)

	orderfile, err := order.NewOrderFileFrom(*orderfilePath)
	if err != nil {
		log.Error().Msgf("Couldn't load orderfile from %s\n due to error: %s", *orderfilePath, err)
		os.Exit(1)
	}

	if *listOrders {
		fmt.Print(strings.Join(orderfile.ListOrdersNames(), "\n"))
		return
	}

	orderName := orderCmd.Arg(0)
	if orderName == "" {
		log.Error().Msg("No order specified, exiting")
		os.Exit(1)
	}

	selectedOrder, ok := orderfile.GetOrder(orderName)
	if !ok {
		log.Error().Msgf(`Couldn't find order "%s"`, orderName)
		log.Info().Msgf("available orders: \n- %s", strings.Join(orderfile.ListOrdersNames(), "\n- "))
		os.Exit(1)
	}

	runner, err := order.NewRunner(
		order.RunnerOptions{
			NoCommand: *noCommands,
		},
	)

	if err != nil {
		log.Error().Msgf("Failed to create runner due to %s", err)
		os.Exit(1)
	}

	if err := runner.RunOrder(selectedOrder); err != nil {
		log.Error().Msgf("Execution of order \"%s\" failed with: \n%s", orderName, err)
		os.Exit(1)
	}
}
