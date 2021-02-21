package cmd

import (
	"github.com/fatih/color"
)

var AsError = color.New(color.FgRed).PrintlnFunc()
var AsInfo = color.New().PrintlnFunc()
var AsWarning = color.New(color.FgYellow).PrintlnFunc()
