package logger

// Level is a log level.
// LevelOff alone turns all logs off.
// LevelAll turns all logs on whenever present in the
// log level caculation. Log levels' bits are or'ed together,
// so, regardless of the order they appear in the caculation,
// if they're present in it, they're turned on.
// For example, to turn only Fatal e Debug levels on,
// the level should be set the following way:
// 	l := logger.New(logger.LevelFatal | logger.LevelDebug, os.Stdout, os.Stderr)
type Level uint8

const (
	LevelOff   Level = 0
	LevelAll   Level = 255
	LevelFatal Level = 1 << iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)
