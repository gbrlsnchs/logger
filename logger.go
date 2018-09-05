package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// LevelOff alone turns all logs off.
//
// LevelAll turns all logs on whenever present in the
// log level caculation.
//
// Log levels' bits are or'ed together,
// so, regardless of the order they appear in the caculation,
// if they're present in it, they're turned on.
//
// For example, to turn only Fatal e Debug levels on,
// the level should be set the following way:
// 	l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelFatal|logger.LevelDebug)
// To turn all levels off:
// 	l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelOff)
// To turn all levels on:
// 	l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelAll)
// All but Debug:
// 	l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelAll^logger.LevelDebug)
const (
	LevelOff uint8 = 1 << iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll uint8 = 254
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
	stdout      io.Writer
	stderr      io.Writer
	flag        int
}

// New creates a new logger according to the log level.
func New(stdout, stderr io.Writer, flag int, lvl uint8) *Logger {
	l := &Logger{
		stdout: stdout,
		stderr: stderr,
		flag:   flag,
	}
	l.SetLevel(lvl)
	return l
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

func (l *Logger) SetLevel(lvl uint8) {
	if lvl&LevelOff > 0 {
		lvl = 0
	}
	l.fatalLogger = build(lvl&LevelFatal, l.stderr, PrefixFatal, l.flag)
	l.errLogger = build(lvl&LevelError, l.stderr, PrefixError, l.flag)
	l.warnLogger = build(lvl&LevelWarn, l.stderr, PrefixWarn, l.flag)
	l.infoLogger = build(lvl&LevelInfo, l.stdout, PrefixInfo, l.flag)
	l.debugLogger = build(lvl&LevelDebug, l.stdout, PrefixDebug, l.flag)
	l.traceLogger = build(lvl&LevelTrace, l.stdout, PrefixTrace, l.flag)
}

// Stderr returns the writer for error logs.
func (l *Logger) Stderr() io.Writer {
	return l.stderr
}

// Stdout returns the writer for ordinary logs.
func (l *Logger) Stdout() io.Writer {
	return l.stdout
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
