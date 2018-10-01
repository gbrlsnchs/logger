package logwrap_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	. "github.com/gbrlsnchs/logwrap"
)

// Values end in a line feed because log package always
// appends one when text doesn't end with line feed.
const (
	TxtFatal = "FATAL"
	TxtError = "ERROR"
	TxtWarn  = "WARN"
	TxtInfo  = "INFO"
	TxtDebug = "DEBUG"
	TxtTrace = "TRACE"
)

var currentLvl int

func TestLogger(t *testing.T) {
	var stderr, stdout strings.Builder
	logger := New(&Options{
		Stderr: &stderr,
		Stdout: &stdout,
		Level:  LevelOff,
	})
	fatalFunc := func(bin string) func(v ...interface{}) {
		return func(v ...interface{}) {
			bin := []byte(bin)
			bin = append([]byte("testdata/"), bin...)
			if os.Getenv("GOOS") == "windows" {
				bin = append(bin, ".exe"...)
			}
			cmd := exec.Command(string(bin), strconv.Itoa(currentLvl), fmt.Sprint(v...))
			cmd.Stderr = &stderr
			cmd.Stdout = &stdout
			cmd.Run()
		}
	}
	testCases := []struct {
		fn     func(v ...interface{})
		txt    string
		stderr string
		stdout string
		lvl    int
	}{
		{logger.Debug, TxtDebug, "", TxtDebug, LevelDebug},
		{formatHelper(logger.Debugf), TxtDebug, "", TxtDebug, LevelDebug},
		{logger.Debugln, TxtDebug, "", TxtDebug, LevelDebug},
		{logger.Error, TxtError, TxtError, "", LevelError},
		{formatHelper(logger.Errorf), TxtError, TxtError, "", LevelError},
		{logger.Errorln, TxtError, TxtError, "", LevelError},
		{logger.Info, TxtInfo, "", TxtInfo, LevelInfo},
		{formatHelper(logger.Infof), TxtInfo, "", TxtInfo, LevelInfo},
		{logger.Infoln, TxtInfo, "", TxtInfo, LevelInfo},
		{logger.Warn, TxtWarn, TxtWarn, "", LevelWarn},
		{formatHelper(logger.Warnf), TxtWarn, TxtWarn, "", LevelWarn},
		{logger.Warnln, TxtWarn, TxtWarn, "", LevelWarn},
		{logger.Trace, TxtTrace, "", TxtTrace, LevelTrace},
		{formatHelper(logger.Tracef), TxtTrace, "", TxtTrace, LevelTrace},
		{logger.Traceln, TxtTrace, "", TxtTrace, LevelTrace},
		{fatalFunc("fatal/fatal"), TxtFatal, TxtFatal, "", LevelFatal},
		{fatalFunc("fatalf/fatalf"), TxtFatal, TxtFatal, "", LevelFatal},
		{fatalFunc("fatalln/fatalln"), TxtFatal, TxtFatal, "", LevelFatal},
	}
	for _, tc := range testCases {
		t.Run(tc.txt, func(t *testing.T) {
			// First, test log level off.
			stdout.Reset()
			stderr.Reset()
			currentLvl = LevelOff
			logger.Reset(&Options{
				Stderr: &stderr,
				Stdout: &stdout,
				Level:  currentLvl,
			})
			tc.fn(tc.txt)
			if want, got := "", strings.TrimRight(stdout.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
			if want, got := "", strings.TrimRight(stderr.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}

			// Then, test all but the specific log level.
			stdout.Reset()
			stderr.Reset()
			currentLvl = LevelAll ^ tc.lvl
			logger.Reset(&Options{
				Stderr: &stderr,
				Stdout: &stdout,
				Level:  currentLvl,
			})
			tc.fn(tc.txt)
			if want, got := "", strings.TrimRight(stdout.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
			if want, got := "", strings.TrimRight(stderr.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}

			// Finally, test only the specific level.
			stdout.Reset()
			stderr.Reset()
			currentLvl = tc.lvl
			logger.Reset(&Options{
				Stderr: &stderr,
				Stdout: &stdout,
				Level:  currentLvl,
			})
			tc.fn(tc.txt)
			if want, got := tc.stdout, strings.TrimRight(stdout.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
			if want, got := tc.stderr, strings.TrimRight(stderr.String(), "\n"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestCaller(t *testing.T) {
	var stderr, stdout strings.Builder
	logger := New(&Options{
		Stderr: &stderr,
		Stdout: &stdout,
		Flag:   log.Lshortfile,
		Level:  LevelDebug,
	})
	logger.Debug("")
	if want, got := "logger_test.go:138: \n", stdout.String(); want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func formatHelper(fn func(string, ...interface{})) func(...interface{}) {
	return func(v ...interface{}) {
		fn("%s", v...)
	}
}
