package logwrap_test

import (
	"log"

	"github.com/gbrlsnchs/logwrap"
)

func Example() {
	logger := logwrap.New(&logwrap.Options{
		Flag:  log.LstdFlags,
		Level: logwrap.LevelAll,
		Prefixes: map[int]string{
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
