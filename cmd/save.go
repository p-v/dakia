package cmd

import (
	"dakia/internal/app"
	"os"
	"strings"
)

func Save(allArgs []string) {
	hasCommand := false
	commandStartIdx := -1
	// Check if the command was provided
	for idx, arg := range allArgs {
		if arg == "--" {
			hasCommand = true
			commandStartIdx = idx
			break
		}
	}

	if !hasCommand {
		AsError("No command passed to save")
		os.Exit(1)
	}

	err := app.StoreCommand(allArgs[1:commandStartIdx], strings.Join(allArgs[commandStartIdx+1:], " "))
	if err != nil {
		AsError("Failed to store command", err)
		os.Exit(1)
	}
	AsInfo("Command stored successfully")
}
