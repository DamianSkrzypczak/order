package order

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"code.sajari.com/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mvdan.cc/sh/v3/interp"
)

func TestNewRunner(t *testing.T) {
	_, err := NewRunner(RunnerOptions{})
	assert.NoError(t, err)
}

func TestNewRunnerErrors(t *testing.T) {
	tmp := defaultOptionsForInternalRunner

	defer func() { defaultOptionsForInternalRunner = tmp }()

	failingOption := func(r *interp.Runner) error { return errors.New("testError") }

	defaultOptionsForInternalRunner = []interp.RunnerOption{
		interp.RunnerOption(failingOption),
	}

	_, err := NewRunner(RunnerOptions{})
	assert.Error(t, err, "testError")
}

func TestRunOrder(t *testing.T) {
	runner, err := NewRunner(RunnerOptions{})
	require.NoError(t, err)

	err = runner.RunOrder(
		&Order{},
	)
	assert.NoError(t, err)
}

func TestRunEmptyOrder(t *testing.T) {
	runner, err := NewRunner(RunnerOptions{})
	require.NoError(t, err)

	err = runner.RunOrder(
		&Order{},
	)
	assert.NoError(t, err)
}

func TestRunOrderWithSyntacticallyWrongCmd(t *testing.T) {
	runner, err := NewRunner(RunnerOptions{NoCommand: true})
	require.NoError(t, err)

	err = runner.RunOrder(
		&Order{
			Script: []Cmd{
				"wr@ng?syn!ax",
			},
		},
	)
	assert.Error(t, err, "exit status", "exit status 127")
}

func mockLoggerStdio(mem storage.FS, filename string) (io.WriteCloser, error) {
	stdout, err := mem.Create(context.Background(), filename)
	if err != nil {
		return nil, err
	}

	output := zerolog.ConsoleWriter{Out: stdout}
	output.FormatLevel = func(i interface{}) string {
		return ""
	}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}
	log.Logger = log.Output(output)

	return stdout, nil
}

func TestRunOrderCommandLogging(t *testing.T) {
	mem := storage.Mem()

	stdout, err := mockLoggerStdio(mem, "testFile")
	require.NoError(t, err)

	runner, err := NewRunner(RunnerOptions{NoCommand: false})
	require.NoError(t, err)

	err = runner.RunOrder(
		&Order{
			Script: []Cmd{
				"(", // syntactically incorrect command
			},
		},
	)
	assert.Error(t, err, "1:1: reached EOF without matching ( with )")

	stdout.Close()

	verifyLoggerWrittenData(t, mem, "testFile", "> (\n")
}

func verifyLoggerWrittenData(t *testing.T, mem storage.FS, filename, expectedContent string) {
	outFile, err := mem.Open(context.Background(), filename)

	defer func() { outFile.Close() }()

	require.NoError(t, err)

	consoleOut, err := ioutil.ReadAll(outFile)
	require.NoError(t, err)

	assert.Equal(t, expectedContent, string(consoleOut))
}
