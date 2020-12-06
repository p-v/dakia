package app

import (
	"os/exec"
	"strings"
)

// ExecuteCommand executes the command
func ExecuteCommand(command string) (output []byte, err error) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	output, err = cmd.Output()

	if err != nil {
		return nil, err
	}

	return output, nil
}
