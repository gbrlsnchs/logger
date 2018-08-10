package logger_test

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

	. "github.com/gbrlsnchs/logger"
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

var currentLvl uint8

func TestLogger(t *testing.T) {
	var stdout, stderr strings.Builder
	l := New(&stdout, &stderr, 0, LevelOff)
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
		lvl    uint8
	}{
		{l.Debug, TxtDebug, TxtDebug, "", LevelDebug},
		{l.Debugln, TxtDebug, TxtDebug, "", LevelDebug},
		{formatHelper(l.Debugf), TxtDebug, TxtDebug, "", LevelDebug},
		{l.Error, TxtError, "", TxtError, LevelError},
		{l.Errorln, TxtError, "", TxtError, LevelError},
		{formatHelper(l.Errorf), TxtError, "", TxtError, LevelError},
		{l.Info, TxtInfo, TxtInfo, "", LevelInfo},
		{l.Infoln, TxtInfo, TxtInfo, "", LevelInfo},
		{formatHelper(l.Infof), TxtInfo, TxtInfo, "", LevelInfo},
		{l.Warn, TxtWarn, "", TxtWarn, LevelWarn},
		{l.Warnln, TxtWarn, "", TxtWarn, LevelWarn},
		{formatHelper(l.Warnf), TxtWarn, "", TxtWarn, LevelWarn},
		{l.Trace, TxtTrace, TxtTrace, "", LevelTrace},
		{l.Traceln, TxtTrace, TxtTrace, "", LevelTrace},
		{formatHelper(l.Tracef), TxtTrace, TxtTrace, "", LevelTrace},
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
			l.SetLevel(currentLvl)
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
			l.SetLevel(currentLvl)
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
			l.SetLevel(currentLvl)
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
	l, txt := fatalHelper(t)
	l.Fatal(txt)
}

func TestFatalf(t *testing.T) {
	l, txt := fatalHelper(t)
	l.Fatalf("%s", txt)
}

func TestFatalln(t *testing.T) {
	l, txt := fatalHelper(t)
	l.Fatalln(txt)
}

func TestCaller(t *testing.T) {
	var stdout, stderr strings.Builder
	l := New(&stdout, &stderr, log.Lshortfile, LevelDebug)
	l.Debug("")
	if want, got := "logger_test.go:142: \n", stdout.String(); want != got {
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

	l := New(os.Stdout, os.Stderr, 0, uint8(lvl))
	return l, txt
}

func formatHelper(fn func(string, ...interface{})) func(...interface{}) {
	return func(v ...interface{}) {
		fn("%s", v...)
	}
}
