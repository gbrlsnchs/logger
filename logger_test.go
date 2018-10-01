package logwrap_test

import (
	"bufio"
	"bytes"
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

var currentLvl Level

func TestLogger(t *testing.T) {
	var stdout, stderr strings.Builder
	logger := New(&Options{
		Stderr: &stderr,
		Stdout: &stdout,
		Level:  LevelOff,
	})
	fatalFunc := func(runFlag string) func(v ...interface{}) {
		return func(v ...interface{}) {
			cmd := exec.Command("go", "test", runFlag, "-short")
			lvl := strconv.Itoa(int(currentLvl))
			env := []string{
				strings.Join([]string{"LOG_LEVEL=", lvl}, ""),
				strings.Join([]string{"TEXT=", fmt.Sprint(v...)}, ""),
			}
			cmd.Env = append(os.Environ(), env...)
			b, _ := cmd.CombinedOutput()

			var buf bytes.Buffer
			buf.Write(b)
			rd := bufio.NewReader(&buf)
			line, _ := rd.ReadBytes('\n')
			if string(line) == "PASS\n" {
				return
			}
			stderr.Write(line)
		}
	}
	testCases := []struct {
		fn     func(v ...interface{})
		txt    string
		stdout string
		stderr string
		lvl    Level
	}{
		{logger.Debug, TxtDebug, TxtDebug, "", LevelDebug},
		{logger.Debugln, TxtDebug, TxtDebug, "", LevelDebug},
		{formatHelper(logger.Debugf), TxtDebug, TxtDebug, "", LevelDebug},
		{logger.Error, TxtError, "", TxtError, LevelError},
		{logger.Errorln, TxtError, "", TxtError, LevelError},
		{formatHelper(logger.Errorf), TxtError, "", TxtError, LevelError},
		{logger.Info, TxtInfo, TxtInfo, "", LevelInfo},
		{logger.Infoln, TxtInfo, TxtInfo, "", LevelInfo},
		{formatHelper(logger.Infof), TxtInfo, TxtInfo, "", LevelInfo},
		{logger.Warn, TxtWarn, "", TxtWarn, LevelWarn},
		{logger.Warnln, TxtWarn, "", TxtWarn, LevelWarn},
		{formatHelper(logger.Warnf), TxtWarn, "", TxtWarn, LevelWarn},
		{logger.Trace, TxtTrace, TxtTrace, "", LevelTrace},
		{logger.Traceln, TxtTrace, TxtTrace, "", LevelTrace},
		{formatHelper(logger.Tracef), TxtTrace, TxtTrace, "", LevelTrace},
		{fatalFunc("-run=TestFatal"), TxtFatal, "", TxtFatal, LevelFatal},
		{fatalFunc("-run=TestFatalf"), TxtFatal, "", TxtFatal, LevelFatal},
		{fatalFunc("-run=TestFatalln"), TxtFatal, "", TxtFatal, LevelFatal},
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

func TestFatal(t *testing.T) {
	logger, txt := fatalHelper(t)
	logger.Fatal(txt)
}

func TestFatalf(t *testing.T) {
	logger, txt := fatalHelper(t)
	logger.Fatalf("%s", txt)
}

func TestFatalln(t *testing.T) {
	logger, txt := fatalHelper(t)
	logger.Fatalln(txt)
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
	if want, got := "logger_test.go:163: \n", stdout.String(); want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func fatalHelper(t *testing.T) (*Logger, string) {
	if !testing.Short() {
		t.SkipNow()
	}
	ll := os.Getenv("LOG_LEVEL")
	txt := os.Getenv("TEXT")
	lvl, _ := strconv.ParseUint(ll, 10, 8)

	logger := New(&Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Level:  Level(lvl),
	})
	return logger, txt
}

func formatHelper(fn func(string, ...interface{})) func(...interface{}) {
	return func(v ...interface{}) {
		fn("%s", v...)
	}
}
