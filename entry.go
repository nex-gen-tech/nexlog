package nexlog

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

// LogEntry is a structure that holds the details of a log entry.
type LogEntry struct {
	logger     *logger         // Reference to the logger creating this entry
	logLevel   Level           // Level of the log
	logTime    time.Time       // Time of the log
	rowMessage string          // The raw log message
	logMessage string          // Formatted log message
	logData    map[string]any  // Additional data to be logged
	caller     *runtime.Frame  // Runtime information about the caller function
	timeStamp  time.Time       // Timestamp of the log
	context    context.Context // Context of the log
	buffer     *bytes.Buffer   // Buffer to hold log message data
}

// newEntry creates a new LogEntry with given logger, logLevel and arguments.
func newEntry(logger *logger, logLevel Level, args ...any) *LogEntry {
	logEntry := &LogEntry{
		logger:    logger,
		logLevel:  logLevel,
		logTime:   time.Now(),
		logData:   map[string]any{},
		timeStamp: time.Now(),
		context:   context.Background(),
		buffer:    &bytes.Buffer{},
	}
	return logEntry
}

// Log logs the provided arguments at the given logLevel.
func (le *LogEntry) Log(logLevel Level, args ...any) {
	le.log(logLevel, args...)
}

// LogF logs the provided arguments at the given logLevel, formatting them as specified.
func (le *LogEntry) LogF(logLevel Level, format string, args ...any) {
	le.log(logLevel, fmt.Sprintf(format, args...))
}

// log is a helper function that logs the provided arguments at the given logLevel.
func (le *LogEntry) log(logLevel Level, args ...any) {
	if le.logger.LogFilter.Filter(logLevel, args...) || le.logger.LogLevel <= logLevel {
		// Set Log Message
		le.logMessage = fmt.Sprint(args...)

		// Set Row Message
		le.rowMessage = fmt.Sprint(args...)

		// Set Log Level
		le.logLevel = logLevel

		// Set Log Message
		le.logMessage = le.logger.LogFormatter.Format(le)

		// Print Log Message
		fmt.Println(le.logMessage)

		// if enableDefaultFileLogHook is true, then write log to file
		if le.logger.enableDefaultFileLogHook {
			le.logger.LogHooks = append(le.logger.LogHooks, newDefaultJsonFileLogHook())
		}

		// Fire Log Hooks
		for _, hook := range le.logger.LogHooks {
			hook.Fire(le)
		}
	}
}

// Info logs the provided arguments at the INFO log level.
func (le *LogEntry) Info(args ...any) {
	le.Log(INF, args...)
}

// InfoF logs the provided arguments at the INFO log level, formatting them as specified.
func (le *LogEntry) InfoF(format string, args ...any) {
	le.LogF(INF, format, args...)
}

// Debug logs the provided arguments at the DEBUG log level.
func (le *LogEntry) Debug(args ...any) {
	le.Log(DBG, args...)
}

// DebugF logs the provided arguments at the DEBUG log level, formatting them as specified.
func (le *LogEntry) DebugF(format string, args ...any) {
	le.LogF(DBG, format, args...)
}

// Warn logs the provided arguments at the WARN log level.
func (le *LogEntry) Warn(args ...any) {
	le.Log(WRN, args...)
}

// WarnF logs the provided arguments at the WARN log level, formatting them as specified.
func (le *LogEntry) WarnF(format string, args ...any) {
	le.LogF(WRN, format, args...)
}

// Error logs the provided arguments at the ERROR log level.
func (le *LogEntry) Error(args ...any) {
	le.Log(ERR, args...)
}

// ErrorF logs the provided arguments at the ERROR log level, formatting them as specified.
func (le *LogEntry) ErrorF(format string, args ...any) {
	le.LogF(ERR, format, args...)
}

// Fatal logs the provided arguments at the FATAL log level. Terminates the program after logging.
func (le *LogEntry) Fatal(args ...any) {
	le.Log(FTL, args...)
	os.Exit(1)
}

// FatalF logs the provided arguments at the FATAL log level, formatting them as specified. Terminates the program after logging.
func (le *LogEntry) FatalF(format string, args ...any) {
	le.LogF(FTL, format, args...)
	os.Exit(1)
}

// WithContext adds a context to the log entry and returns the updated LogEntry.
func (le *LogEntry) WithContext(ctx context.Context) *LogEntry {
	le.context = ctx
	return le
}

// WithField adds a key-value pair to the log entry and returns the updated LogEntry.
func (le *LogEntry) WithField(key string, value any) *LogEntry {
	le.logData[key] = value
	return le
}

// WithFields adds multiple key-value pairs to the log entry and returns the updated LogEntry.
func (le *LogEntry) WithFields(data map[string]any) *LogEntry {
	for key, value := range data {
		le.logData[key] = value
	}

	return le
}

// WithTime sets the log time to the provided time and returns the updated LogEntry.
func (le *LogEntry) WithTime(t time.Time) *LogEntry {
	le.logTime = t
	return le
}

// WithIdentifier sets the logger's identifier to the provided string and returns the updated LogEntry.
func (le *LogEntry) WithIdentifier(identifier string) *LogEntry {
	le.logger.Ident = identifier
	return le
}
