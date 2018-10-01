package logwrap

import (
	"io"
	"log"
)

type Options struct {
	Stderr   io.Writer
	Stdout   io.Writer
	Flag     int
	Level    Level
	Prefixes map[Level]string
}

func (opt *Options) build(lvl Level) *log.Logger {
	if opt.Level&LevelOff == 0 && opt.Level&lvl > 0 {
		w := opt.Stderr
		if lvl >= LevelInfo {
			w = opt.Stdout
		}
		return log.New(w, opt.Prefixes[lvl], opt.Flag)
	}
	return nil
}
