package logger

import "github.com/fatih/color"

var (
	PrefixFatal = color.New(color.FgWhite).Sprint("FATAL: ")
	PrefixError = color.New(color.FgRed).Sprint("ERROR: ")
	PrefixWarn  = color.New(color.FgYellow).Sprint("WARN: ")
	PrefixInfo  = color.New(color.FgGreen).Sprint("INFO: ")
	PrefixDebug = color.New(color.FgCyan).Sprint("DEBUG: ")
	PrefixTrace = color.New(color.FgBlue).Sprint("TRACE: ")
)
