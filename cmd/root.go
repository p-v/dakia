package cmd

import (
	"dakia/internal/app"
	"fmt"
	"os"
	"strings"
)

// Execute the command
func Execute() {
	allArgs := os.Args[1:]

	// TODO add more validations
	if len(allArgs) == 0 {
		fmt.Println("No args passed")
		os.Exit(1)
	}

	isSaveMode := allArgs[0] == "-s"
	isFindMode := allArgs[0] == "-f"

	if isSaveMode {
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
			fmt.Println("No command passed to save")
			os.Exit(1)
		}

		err := app.StoreCommand(allArgs[1:commandStartIdx], strings.Join(allArgs[commandStartIdx+1:], " "))
		if err != nil {
			fmt.Println("Failed to store command", err)
			os.Exit(1)
		}
		fmt.Println("Command stored successfully")
	} else if isFindMode {
		commandDetail, err := app.ListCommands(allArgs[1:])
		if err != nil {
			fmt.Println("Failed to list commands", err)
			os.Exit(1)
		}

		commands := commandDetail.Commands
		executable := commandDetail.Executable

		if executable != "" {
			fmt.Println("Execuctable: ", executable)
			fmt.Println("===================")
		}

		fmt.Println("Available Commands:")
		fmt.Println("===================")
		for _, command := range commands {
			fmt.Println(command)
		}
	} else {
		// Happy case
		command, err := app.GetCommand(allArgs[:])
		if err != nil {
			fmt.Println("Failed to retrieve command", err)
			os.Exit(1)
		}
		output, err := app.ExecuteCommand(string(command))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(string(output))
	}
}
