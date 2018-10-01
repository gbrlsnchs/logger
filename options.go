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

func (o *Options) build(lvl Level) *log.Logger {
	if o.Level&LevelOff == 0 && o.Level&lvl > 0 {
		w := o.Stderr
		if lvl >= LevelInfo {
			w = o.Stdout
		}
		return log.New(w, o.Prefixes[lvl], o.Flag)
	}
	return nil
}
