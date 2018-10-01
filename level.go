package logwrap

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
	LevelOff = 1 << iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll = LevelFatal | LevelError | LevelWarn | LevelInfo | LevelDebug | LevelTrace
)
