package logger

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
// 	l := logger.New(logger.LevelFatal | logger.LevelDebug, os.Stdout, os.Stderr)
// To turn all levels off:
// 	l := logger.New(logger.LevelOff, os.Stdout, os.Stderr)
// And to turn all on:
// 	l := logger.New(logger.LevelAll, os.Stdout, os.Stderr)
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
