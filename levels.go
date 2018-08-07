package logger

type Level int8

const (
	LevelOff   Level = 0
	LevelFatal       = 1 << iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)
