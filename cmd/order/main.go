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
	hideCommands := orderCmd.BoolP("no-commands", "n", false, "do not show currently executed command")
	orderfilePath := orderCmd.StringP("path", "p", "Orderfile.yml", "path to orderfile")

	err := orderCmd.Parse(os.Args[1:])

	if err == pflag.ErrHelp {
		return
	}

	setupLogger(*debugModeOn)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	orderfile, err := order.NewOrderFileFrom(*orderfilePath)
	if err != nil {
		log.Error().Msgf("Couldn't load orderfile from %s\n", *orderfilePath)
		os.Exit(1)
	}

	if *listOrders {
		fmt.Print(strings.Join(orderfile.ListOrdersNames(), "\n"))
		return
	}

	if orderNameFromCLI := orderCmd.Arg(0); orderNameFromCLI != "" {
		for orderNameFromFile, order := range orderfile.Orders {
			if orderNameFromCLI == orderNameFromFile {
				if err := order.Run(*hideCommands); err != nil {
					log.Error().Msgf("Execution of order \"%s\" failed with: \n%s", orderNameFromFile, err)
				}

				return
			}
		}

		log.Error().Msgf(`Couldn't find order "%s"`, orderNameFromCLI)
		log.Info().Msgf("available orders: \n- %s", strings.Join(orderfile.ListOrdersNames(), "\n- "))
		os.Exit(1)
	}

	log.Error().Msg("No order specified, exiting")
	os.Exit(1)
}
