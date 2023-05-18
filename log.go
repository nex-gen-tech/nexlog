package nexlog

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// Create a Custom GoLang Logger package called `nexlog` that implements the `Logger` interface. The `nexlog` package should have the following features, SOLID principles and all the best practices of GoLang:

// 1) It should have a plugable LogFormatter interface that can be used to format the log output. You have to create two basic formatters: `JSONFormatter` and `TextFormatter`. The `TextFormatter` should be the default formatter.
// 2) It should have a plugable LogWriter interface that can be used to write the log output. You have to create two basic writers: `FileWriter` and `StdoutWriter`. The `StdoutWriter` should be the default writer.
// 3) It should have a plugable LogLevel interface that can be used to set the log level. You have to create three basic log levels: `INF`, `DBG` and `ERR`. The `INF` should be the default log level.
// 4) It should have a plugable LogFilter interface that can be used to filter the log output. You have to create two basic filters: `NoFilter` and `LevelFilter`. The `NoFilter` should be the default filter.
// 5) it should also have the hooking feature that can be used to hook the log output to a third party service. There should a LogHook interface and a AddHook function which should be used to add a hook to the logger. You have to create two basic hooks: `SlackHook` and `EmailHook`. The `SlackHook` should be the default hook.
// 6) It should have a `NewLogger` function that can be used to create a new logger instance. The `NewLogger` function should take the following arguments:
//     - `logLevel`: The log level to be used by the logger. The default value should be `INFO`.
//     - `logFormatter`: The log formatter to be used by the logger. The default value should be `TextFormatter`.
//     - `logWriter`: The log writer to be used by the logger. The default value should be `StdoutWriter`.
//     - `logFilter`: The log filter to be used by the logger. The default value should be `NoFilter`.
//     - `logHook`: The log hook to be used by the logger. The default value should be `SlackHook`.
// 7) It should have the following methods:
//     - `Debug`: This method should be used to log the debug messages.
//     - `Info`: This method should be used to log the info messages.
//     - `Error`: This method should be used to log the error messages.
//     - `SetLogLevel`: This method should be used to set the log level.
//     - `SetLogFormatter`: This method should be used to set the log formatter.
//     - `SetLogWriter`: This method should be used to set the log writer.
//     - `SetLogFilter`: This method should be used to set the log filter.
//     - `SetLogHook`: This method should be used to set the log hook.
//     - `AddHook`: This method should be used to add a hook to the logger.
// 8) It should have the following constants:
//     - `DBG`
//     - `INF`
//     - `ERR`
//     - `WRN`
//     - `FTL`
// 9) It should have the following variables:
//     - `logLevel`
//     - `logFormatter`
//     - `logWriter`
//     - `logFilter`
//     - `logHook`
// 10) It should use the following packages:
// - log - https://golang.org/pkg/log/
// - color - github.com/fatih/color - for colorizing the log output based on the log level
//  - pool - sync.Pool
//  - json - encoding/json
// - time - time
// - os - os
// - io - io
// - bytes - bytes
// - fmt - fmt
// - errors - errors
// - strings - strings
// - reflect - reflect
// - sync - sync
// - net/smtp - net/smtp
// - github.com/slack-go/slack - github.com/slack-go/slack
// - github.com/getsentry/sentry-go - github.com/getsentry/sentry-go
// - github.com/lestrrat-go/file-rotatelogs - github.com/lestrrat-go/file-rotatelogs
// - github.com/mattn/go-colorable - github.com/mattn/go-colorable
// - github.com/mattn/go-isatty - github.com/mattn/go-isatty

type Level int

const (
	DBG Level = iota
	INF
	ERR
	WRN
	FTL
)

// String - String function
func (l Level) String() string {
	switch l {
	case DBG:
		return "DBG"
	case INF:
		return "INF"
	case ERR:
		return "ERR"
	case WRN:
		return "WRN"
	case FTL:
		return "FTL"
	default:
		return "INF"
	}
}

type Logger interface {
	Debug(args ...any)
	DebugF(format string, args ...any)
	Info(args ...any)
	InfoF(format string, args ...any)
	Warn(args ...any)
	WarnF(format string, args ...any)
	Error(args ...any)
	ErrorF(format string, args ...any)
	Log(logLevel Level, args ...any)
	LogF(logLevel Level, format string, args ...any)
	// SetLogLevel(logLevel Level)
	// SetLogFormatter(logFormatter LogFormatter)
	// SetLogFilter(logFilter LogFilter)
	// AddHook(logHook LogHook)
}

// logger - logger struct
type logger struct {
	LogLevel     Level
	LogFormatter LogFormatter
	LogFilter    LogFilter
	LogHook      []LogHook
	Ident        string
	EntryPool    sync.Pool
}

// NewLogger - NewLogger function
func NewLogger(ident string) Logger {
	return &logger{
		LogLevel:     INF,
		LogFormatter: NewDefaultTextFormatter(),
		LogFilter:    NewDefaultNoFilter(),
		LogHook:      []LogHook{NewDefaultSlackHook()},
		Ident:        ident,
	}
}

// Log - Log function prints the log message
func (l *logger) Log(logLevel Level, args ...any) {
	var entry *LogEntry
	entry, ok := l.EntryPool.Get().(*LogEntry)
	if !ok {
		entry = NewEntry(l, logLevel, args...)
	}
	entry.caller = getCaller()
	entry.Log(logLevel, args...)
	l.EntryPool.Put(entry)
}

// LogF - LogF function prints the log message with the given format
func (l *logger) LogF(logLevel Level, format string, args ...any) {
	var entry *LogEntry
	entry, ok := l.EntryPool.Get().(*LogEntry)
	if !ok {
		entry = NewEntry(l, logLevel, args...)
	}
	entry.caller = getCaller()
	entry.Log(logLevel, fmt.Sprintf(format, args...))
	l.EntryPool.Put(entry)
}

// Debug - Debug function
func (l *logger) Debug(args ...any) {
	l.Log(DBG, args...)
}

// DebugF - DebugF function\
func (l *logger) DebugF(format string, args ...any) {
	l.LogF(DBG, format, args...)
}

// Info - Info function
func (l *logger) Info(args ...any) {
	l.Log(INF, args...)
}

// InfoF - InfoF function
func (l *logger) InfoF(format string, args ...any) {
	l.LogF(INF, format, args...)
}

// Warn - Warn function
func (l *logger) Warn(args ...any) {
	l.Log(WRN, args...)
}

// WarnF - WarnF function
func (l *logger) WarnF(format string, args ...any) {
	l.LogF(WRN, format, args...)
}

// Error - Error function
func (l *logger) Error(args ...any) {
	l.Log(ERR, args...)
}

// ErrorF - ErrorF function
func (l *logger) ErrorF(format string, args ...any) {
	l.LogF(ERR, format, args...)
}

// getCaller - will return the caller
// func getCaller() *runtime.Frame {
// 	pc := make([]uintptr, 10)
// 	runtime.Callers(4, pc)
// 	frames := runtime.CallersFrames(pc)
// 	frame, _ := frames.Next()
// 	return &frame
// }

func getCaller() *runtime.Frame {
	var frame *runtime.Frame
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])

	for {
		f, more := frames.Next()
		if !more {
			break
		}

		// skip over frames from log and runtime package
		if strings.Contains(f.Function, "nexlog") || strings.Contains(f.Function, "runtime") {
			continue
		}

		frame = &f
		break
	}
	return frame
}

// func getPackageName(f string) string {
// 	for {
// 		lastPeriod := strings.LastIndex(f, ".")
// 		lastSlash := strings.LastIndex(f, "/")
// 		if lastPeriod > lastSlash {
// 			f = f[:lastPeriod]
// 		} else {
// 			break
// 		}
// 	}

// 	return f
// }
