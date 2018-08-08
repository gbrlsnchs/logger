package logger

import (
	"io"
	"log"

	"github.com/fatih/color"
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
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelFatal|logger.LevelDebug)
// To turn all levels off:
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelOff)
// To turn all levels on:
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelAll)
// All but Debug:
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelAll^logger.LevelDebug)
const (
	LevelOff   uint8 = 0
	LevelAll   uint8 = 255
	LevelFatal uint8 = 1 << iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

var Flag = log.Ldate | log.Lmicroseconds | log.Lshortfile

type Logger struct {
	fatalLogger *log.Logger
	errLogger   *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	debugLogger *log.Logger
	traceLogger *log.Logger
}

// New creates a new logger according to the log level.
func New(stdout, stderr io.Writer, lvl uint8) *Logger {
	l := &Logger{}
	if lvl == LevelOff {
		return l
	}

	if lvl&LevelFatal > 0 {
		l.fatalLogger = log.New(stderr, PrefixFatal, Flag)
	}
	if lvl&LevelError > 0 {
		l.errLogger = log.New(stderr, PrefixError, Flag)
	}
	if lvl&LevelWarn > 0 {
		l.warnLogger = log.New(stderr, PrefixWarn, Flag)
	}
	if lvl&LevelInfo > 0 {
		l.infoLogger = log.New(stdout, PrefixInfo, Flag)
	}
	if lvl&LevelDebug > 0 {
		l.debugLogger = log.New(stdout, PrefixDebug, Flag)
	}
	if lvl&LevelTrace > 0 {
		l.traceLogger = log.New(stdout, PrefixTrace, Flag)
	}
	return l
}

func (l *Logger) Debug(v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Print(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Printf(format, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.debugLogger == nil {
		return
	}
	l.debugLogger.Println(v...)
}

func (l *Logger) DisableColor() {
	color.NoColor = true
}

func (l *Logger) Error(v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Print(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Printf(format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	if l.errLogger == nil {
		return
	}
	l.errLogger.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Fatal(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Fatalf(format, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	if l.fatalLogger == nil {
		return
	}
	l.fatalLogger.Fatalln(v...)
}

func (l *Logger) Info(v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Print(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Printf(format, v...)
}

func (l *Logger) Infoln(format string, v ...interface{}) {
	if l.infoLogger == nil {
		return
	}
	l.infoLogger.Println(v...)
}

func (l *Logger) Trace(v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Print(v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Printf(format, v...)
}

func (l *Logger) Traceln(v ...interface{}) {
	if l.traceLogger == nil {
		return
	}
	l.traceLogger.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Print(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Printf(format, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.warnLogger == nil {
		return
	}
	l.warnLogger.Println(v...)
}
