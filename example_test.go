package logwrap_test

import (
	"log"
	"os"

	"github.com/gbrlsnchs/logwrap"
)

func Example() {
	logger := logwrap.New(&logwrap.Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Flag:   log.LstdFlags,
		Level:  logwrap.LevelAll,
		Prefixes: map[logwrap.Level]string{
			logwrap.LevelFatal: "FATAL::",
			logwrap.LevelError: "ERROR::",
			logwrap.LevelWarn:  "WARN::",
			logwrap.LevelInfo:  "INFO::",
			logwrap.LevelDebug: "DEBUG::",
			logwrap.LevelTrace: "TRACE::",
		},
	})
	logger.Debug("Hello World!")
}
