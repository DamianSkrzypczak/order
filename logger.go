package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func newLogger(debug, noLogLevel, noColor bool) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.NoColor = noColor
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}

	if noLogLevel {
		output.FormatLevel = func(i interface{}) string {
			return ""
		}
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return log.Output(output)
}

func newHelpLogger() zerolog.Logger {
	return newLogger(false, true, true)
}
