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
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelFatal|logger.LevelDebug)
// To turn all levels off:
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelOff)
// And to turn all on:
// 	l := logger.New(os.Stdout, os.Stderr, logger.LevelAll)
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
