package cmd

import (
	"dakia/internal/app"
	"os"
)

func Find(allArgs []string) {
	commandDetail, err := app.ListCommands(allArgs[1:])
	if err != nil {
		AsError("Failed to list commands", err)
		os.Exit(1)
	}

	commands := commandDetail.Commands
	executable := commandDetail.Executable

	if executable != "" {
		AsInfo("Execuctable: ", executable)
		AsInfo("===================")
	}

	AsInfo("Available Commands:")
	AsInfo("===================")
	for _, command := range commands {
		AsInfo(command)
	}

}
