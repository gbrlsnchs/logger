package main

import (
	"os"
	"strconv"

	"github.com/gbrlsnchs/logwrap"
)

func main() {
	lvl, _ := strconv.Atoi(os.Args[1])
	logger := logwrap.New(&logwrap.Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Level:  lvl,
	})
	logger.Fatalf("%s", os.Args[2])
}
