package logwrap

import (
	"fmt"
	"log"
	"os"
)

var (
	PrefixFatal string
	PrefixError string
	PrefixWarn  string
	PrefixInfo  string
	PrefixDebug string
	PrefixTrace string
)

type Logger struct {
	fatalLogger *log.Logger
	errLogger   *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	debugLogger *log.Logger
	traceLogger *log.Logger
}

// New creates a new logger according to the log level.
func New(opt *Options) *Logger {
	var logger Logger
	logger.Reset(opt)
	return &logger
}

func (l *Logger) Debug(v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorln(v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Info(v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Infoln(v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Reset(opt *Options) {
	if opt.Level == 0 {
		opt.Level = LevelOff
	}
	l.fatalLogger = opt.build(LevelFatal)
	l.errLogger = opt.build(LevelError)
	l.warnLogger = opt.build(LevelWarn)
	l.infoLogger = opt.build(LevelInfo)
	l.debugLogger = opt.build(LevelDebug)
	l.traceLogger = opt.build(LevelTrace)
}

func (l *Logger) Trace(v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Traceln(v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logger) Warn(v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Output(2, fmt.Sprintln(v...))
}
