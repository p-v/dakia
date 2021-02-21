package executor

import (
	"dakia/internal/app"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	output, err := app.ExecuteCommand("echo hello")
	if err != nil {
		t.Log("Command returned an error")
		t.Fail()
	}

	if string(output) != "hello\n" {
		t.Logf("Output (%s) is not same as expected (%s)", output, "hello")
		t.Fail()
	}
}

func TestExecuteCommandWithOptionalArgs(t *testing.T) {
	output, err := app.ExecuteCommand("echo -n hello")
	if err != nil {
		t.Log("Command returned an error")
		t.Fail()
	}

	if string(output) != "hello" {
		t.Logf("Output (%s) is not same as expected (%s)", output, "hello")
		t.Fail()
	}
}
