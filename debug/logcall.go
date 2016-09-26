package debug

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Default logger
var defaultLogger DebugLogger

type DebugLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

func init() {
	defaultLogger = log.New(os.Stderr, "", log.LstdFlags)
}

func SetDefaultLogger(dl DebugLogger) {
	defaultLogger = dl
}

var LogCallEnabled bool = false

func getCallerName(i int) string {
	pc, _, _, ok := runtime.Caller(i)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	return me.Name()
}

func Fn() string {
	return getCallerName(2)
}

func LogCall() {
	if LogCallEnabled {
		defaultLogger.Print(getCallerName(2))
	}
}

func Log(msgAndArgs ...interface{}) {
	if len(msgAndArgs) == 0 {
		return
	}
	if LogCallEnabled {
		msgAndArgs[0] = fmt.Sprintf("%v %v", getCallerName(2), msgAndArgs[0])
	}
	defaultLogger.Print(msgAndArgs...)
}

func Logf(msgAndArgs ...interface{}) {
	if len(msgAndArgs) == 0 {
		return
	}
	if LogCallEnabled {
		msgAndArgs[0] = fmt.Sprintf("%v %v", getCallerName(2), msgAndArgs[0])
	}
	defaultLogger.Printf(msgAndArgs[0].(string), msgAndArgs[1:]...)
}
