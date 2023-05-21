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
	mu                       sync.Mutex
}

type Config struct {
	LogLevel       Level
	LogFormatter   LogFormatter
	LogFilter      LogFilter
	EnableCaller   bool
	EnableFileHook bool
}

// New - New function
func New(ident string, config ...Config) Logger {
	l := logger{
		LogLevel:                 INF,
		LogFormatter:             newDefaultTextFormatter(),
		LogFilter:                newDefaultNoFilter(),
		LogHooks:                 []LogHook{},
		Ident:                    ident,
		isCaller:                 false,
		enableDefaultFileLogHook: false,
	}

	if len(config) > 0 {
		// Set log level if provided
		if config[0].LogLevel != 0 {
			l.LogLevel = config[0].LogLevel
		}
		// Set log formatter if provided
		if config[0].LogFormatter != nil {
			l.LogFormatter = config[0].LogFormatter
		}
		// Set log filter if provided
		if config[0].LogFilter != nil {
			l.LogFilter = config[0].LogFilter
		}
		// Enable caller if provided
		if config[0].EnableCaller {
			l.EnableCaller()
		}
		// Enable default file log hook if provided
		if config[0].EnableFileHook {
			l.EnableDefaultFileLogHook()
		}
		// Add default file log hook if enabled
		if l.enableDefaultFileLogHook {
			l.AddHook(newDefaultJsonFileLogHook())
		}
	}

	// Add entry pool
	l.EntryPool.New = func() any {
		return newEntry(&l, INF)
	}

	// return logger
	return &l
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
	l.mu.Lock()
	defer l.mu.Unlock()
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
	l.mu.Lock()
	defer l.mu.Unlock()
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
