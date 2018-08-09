package logger

import (
	"io"
	"log"
)

func build(lvl uint8, w io.Writer, prefix string, flag int) *log.Logger {
	if lvl > 0 {
		return log.New(w, prefix, flag)
	}
	return nil
}
