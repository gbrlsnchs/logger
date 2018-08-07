package logger

import (
	"io"
	"log"
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
func New(lvl Level, stdout, stderr io.Writer) *Logger {
	// TODO: color logs.
	l := &Logger{}
	flag := log.Ldate | log.Lmicroseconds | log.Lshortfile
	if lvl == LevelOff {
		return l
	}
	if lvl&LevelFatal > 0 {
		l.fatalLogger = log.New(stderr, "FATAL: ", flag)
	}
	if lvl&LevelError > 0 {
		l.errLogger = log.New(stderr, "ERROR: ", flag)
	}
	if lvl&LevelWarn > 0 {
		l.warnLogger = log.New(stderr, "WARNING: ", flag)
	}
	if lvl&LevelInfo > 0 {
		l.infoLogger = log.New(stdout, "INFO: ", flag)
	}
	if lvl&LevelDebug > 0 {
		l.debugLogger = log.New(stdout, "DEBUG: ", flag)
	}
	if lvl&LevelTrace > 0 {
		l.traceLogger = log.New(stdout, "TRACE: ", flag)
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
