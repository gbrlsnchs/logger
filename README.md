# logwrap (Log wrapper for Go)
[![Build Status](https://travis-ci.org/gbrlsnchs/logwrap.svg?branch=master)](https://travis-ci.org/gbrlsnchs/logwrap)
[![Build status](https://ci.appveyor.com/api/projects/status/ekck6k62bmrpdl8c/branch/master?svg=true)](https://ci.appveyor.com/project/gbrlsnchs/logwrap/branch/master)
[![Sourcegraph](https://sourcegraph.com/github.com/gbrlsnchs/logwrap/-/badge.svg)](https://sourcegraph.com/github.com/gbrlsnchs/logwrap?badge)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/logwrap?status.svg)](https://godoc.org/github.com/gbrlsnchs/logwrap)

## About
This package is a log wrapper that supports the same log levels from [Log4j](https://logging.apache.org/log4j/).  
However, it doesn't impose a hierarchy, allowing levels to be cherry-picked through [bitwise operations](https://en.wikipedia.org/wiki/Bitwise_operation).

| Level | Variable            | Value      |
| ----- | ------------------- |:----------:|
| OFF   | `logwrap.LevelOff`   | `00000001` |
| FATAL | `logwrap.LevelFatal` | `00000010` |
| ERROR | `logwrap.LevelError` | `00000100` |
| WARN  | `logwrap.LevelWarn`  | `00001000` |
| INFO  | `logwrap.LevelInfo`  | `00010000` |
| DEBUG | `logwrap.LevelDebug` | `00100000` |
| TRACE | `logwrap.LevelTrace` | `01000000` |
| ALL   | `logwrap.LevelAll`   | `11111111` |

## Usage
Full documentation [here](https://godoc.org/github.com/gbrlsnchs/logwrap).

### Installing
#### Go 1.10
`vgo get -u github.com/gbrlsnchs/logwrap`
#### Go 1.11
`go get -u github.com/gbrlsnchs/logwrap`

### Importing
```go
import (
	// ...

	"github.com/gbrlsnchs/logwrap"
)
```

### Setting all levels on
```go
logger := logwrap.New(&logwrap.Options{
	Stderr: os.Stderr,
	Stdout: os.Stdout,
	Flag:   log.LStdFlags,
	Level:  logwrap.LevelAll,
})
logger.Trace("Hello World!")
logger.Debug("Hello World!")
logger.Info("Hello World!")
logger.Warn("Hello World!")
logger.Error("Hello World!")
logger.Fatal("Hello World!")
```

### Cherry-picking levels
```go
logger := logwrap.New(&logwrap.Options{
	Stderr: os.Stderr,
	Stdout: os.Stdout,
	Flag:   log.LstdFlags,
	Level:  logwrap.LevelError | logwrap.LevelDebug // only Error and Debug levels will print
})
logger.Info("this will not be printed")
```

### Excluding a level
```go
logger := logwrap.New(&logwrap.Options{
	Stderr: os.Stderr,
	Stdout: os.Stdout,
	Flag:   log.LstdFlags,
	Level:  logwrap.LevelAll ^ logwrap.LevelFatal) // all levels but Fatal will print
logger.Fatal("this will neither be printed nor exit the process")
```

### Setting custom prefixes
```go
logger := logwrap.New(&logwrap.Options{
	Stderr: os.Stderr,
	Stdout: os.Stdout,
	Flag:   log.LstdFlags,
	Level:  logwrap.LevelAll,
	Prefixes: map[logwrap.Level]string{
		logwrap.LevelFatal: "FATAL::",
		logwrap.LevelError: "ERROR::",
		logwrap.LevelWarn:  "WARN::",
		logwrap.LevelInfo:  "INFO::",
		logwrap.LevelDebug: "DEBUG::",
		logwrap.LevelTrace: "TRACE::",
	},
})
logger.Debug("Hello World!")
```

## Contributing
### How to help
- For bugs and opinions, please [open an issue](https://github.com/gbrlsnchs/logwrap/issues/new)
- For pushing changes, please [open a pull request](https://github.com/gbrlsnchs/logwrap/compare)
