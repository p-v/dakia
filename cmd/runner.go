package cmd

import (
	"dakia/internal/app"
	"os"
)

func Run(allArgs []string) {
	// Happy case
	command, err := app.GetCommand(allArgs[:])
	if err != nil {
		AsError("Failed to retrieve command", err)
		os.Exit(1)
	}
	output, err := app.ExecuteCommand(string(command))
	if err != nil {
		AsError(err.Error())
		os.Exit(1)
	}

	AsInfo(string(output))
}
