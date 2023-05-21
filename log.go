package nexlog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/gookit/color"
)

// Level - Level type is a Enum for log levels
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

// Color - Color function
func (l Level) Color() string {
	switch l {
	case DBG:
		return color.BgYellow.Render(fmt.Sprintf(" %s ", l.String()))
	case INF:
		return color.BgHiGreen.Render(fmt.Sprintf(" %s ", l.String()))
	case ERR:
		return color.BgRed.Render(fmt.Sprintf(" %s ", l.String()))
	case WRN:
		return color.BgCyan.Render(fmt.Sprintf(" %s ", l.String()))
	case FTL:
		return color.BgHiRed.Render(fmt.Sprintf(" %s ", l.String()))
	default:
		return color.BgHiGreen.Render(fmt.Sprintf(" %s ", l.String()))
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
	Fatal(args ...any)
	FatalF(format string, args ...any)
	Log(logLevel Level, args ...any)
	LogF(logLevel Level, format string, args ...any)
	SetLogLevel(logLevel Level)
	SetLogFormatter(logFormatter LogFormatter)
	SetLogFilter(logFilter LogFilter)
	AddHook(logHook LogHook)
	EnableCaller()
	EnableDefaultFileLogHook()
}

// logger - logger struct
type logger struct {
	LogLevel                 Level
	LogFormatter             LogFormatter
	LogFilter                LogFilter
	LogHooks                 []LogHook
	Ident                    string
	EntryPool                sync.Pool
	isCaller                 bool
	enableDefaultFileLogHook bool
}

// New - New function
func New(ident string) Logger {
	return &logger{
		LogLevel:                 INF,
		LogFormatter:             newDefaultTextFormatter(),
		LogFilter:                newDefaultNoFilter(),
		LogHooks:                 []LogHook{},
		Ident:                    ident,
		isCaller:                 false,
		enableDefaultFileLogHook: false,
	}
}

// EnableCaller  - enable caller info
func (l *logger) EnableCaller() {
	l.isCaller = true
}

// EnableDefaultFileLogHook - enable default file log hook
func (l *logger) EnableDefaultFileLogHook() {
	l.enableDefaultFileLogHook = true
}

// Log - Log function prints the log message
func (l *logger) Log(logLevel Level, args ...any) {
	var entry *LogEntry
	entry, ok := l.EntryPool.Get().(*LogEntry)
	if !ok {
		entry = newEntry(l, logLevel, args...)
	}
	if l.isCaller {
		entry.caller = getCaller()
	}
	entry.Log(logLevel, args...)
	l.EntryPool.Put(entry)
}

// LogF - LogF function prints the log message with the given format
func (l *logger) LogF(logLevel Level, format string, args ...any) {
	var entry *LogEntry
	entry, ok := l.EntryPool.Get().(*LogEntry)
	if !ok {
		entry = newEntry(l, logLevel, args...)
	}
	if l.isCaller {
		entry.caller = getCaller()
	}
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

// getCaller - getCaller function returns the caller of the function
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

// SetLogLevel - SetLogLevel function sets the log level
func (l *logger) SetLogLevel(logLevel Level) {
	l.LogLevel = logLevel
}

// SetLogFormatter - SetLogFormatter function sets the log formatter
func (l *logger) SetLogFormatter(logFormatter LogFormatter) {
	l.LogFormatter = logFormatter
}

// SetLogFilter - SetLogFilter function  sets the log filter
func (l *logger) SetLogFilter(logFilter LogFilter) {
	l.LogFilter = logFilter
}

// AddHook - AddHook function adds a log hook to the logger
func (l *logger) AddHook(logHook LogHook) {
	l.LogHooks = append(l.LogHooks, logHook)
}

// Fatal - Fatal function
func (l *logger) Fatal(args ...any) {
	l.Log(FTL, args...)
	os.Exit(1)
}

// FatalF - FatalF function
func (l *logger) FatalF(format string, args ...any) {
	l.LogF(FTL, format, args...)
	os.Exit(1)
}
