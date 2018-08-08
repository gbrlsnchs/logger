package logger

type Level uint8

const (
	LevelOff   Level = 0
	LevelFatal Level = 1 << iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll Level = 255
)
