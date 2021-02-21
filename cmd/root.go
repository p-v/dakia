package cmd

import (
	"os"
)

func showHelp() {
	AsInfo("Usage: dakia -s | -f | -h")
	AsInfo("-s, Used for saving a command")
	AsInfo("-f, Used for finding a command")
	AsInfo("-h, Show help")
}

// Execute the command
func Execute() {
	allArgs := os.Args[1:]

	if len(allArgs) == 0 {
		AsWarning("No args passed")
		os.Exit(1)
	}

	switch allArgs[0] {
	case "-s":
		Save(allArgs)
	case "-f":
		Find(allArgs)
	case "-h":
		showHelp()
	default:
		Run(allArgs)
	}
}
