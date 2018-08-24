# logger (filtered-by-level logger for Go)
[![Build Status](https://travis-ci.org/gbrlsnchs/logger.svg?branch=master)](https://travis-ci.org/gbrlsnchs/logger)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/logger?status.svg)](https://godoc.org/github.com/gbrlsnchs/logger)

## About
This package is a logger with level filtering.

### Supported levels:
| Level | Enum                | Value      |
| ----- | ------------------- |:----------:|
| OFF   | `logger.LevelOff`   | `00000001` |
| FATAL | `logger.LevelFatal` | `00000010` |
| ERROR | `logger.LevelError` | `00000100` |
| WARN  | `logger.LevelWarn`  | `00001000` |
| INFO  | `logger.LevelInfo`  | `00010000` |
| DEBUG | `logger.LevelDebug` | `00100000` |
| TRACE | `logger.LevelTrace` | `01000000` |
| ALL   | `logger.LevelAll`   | `11111111` |

## Usage
Full documentation [here].

## Example
### Simple usage
```go
l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelAll)
l.Trace("Hello", "World!")
l.Debug("Hello World!")
l.Info("Hello World!")
l.Warn("Hello World!")
l.Error("Hello World!")
l.Fatal("Hello World!")
```

### Filter levels
```go
l := logger.New(os.Stdout, os.Stderr, log.LstdFlags, logger.LevelError|logger.LevelDebug) // only Error and Debug levels will print
l.Info("this will not be printed")

l.SetLevel(logger.LevelAll^logger.LevelFatal) // all levels but Fatal
l.Fatal("this will not exit the process")
```

## Contribution
### How to help:
- Pull Requests
- Issues
- Opinions

[here]: https://godoc.org/github.com/gbrlsnchs/logger
