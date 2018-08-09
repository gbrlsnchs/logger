package logger_test

import (
	"log"
	"os"

	"github.com/gbrlsnchs/logger"
)

func Example() {
	l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelAll)
	l.Debug("Hello World!")
}
