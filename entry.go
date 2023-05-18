package nexlog

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"time"
)

// LogEntry - LogEntry struct
type LogEntry struct {
	logger     *logger
	logLevel   Level
	logTime    time.Time
	logMessage string
	logData    map[string]any
	caller     *runtime.Frame
	timeStamp  time.Time
	context    context.Context
	buffer     *bytes.Buffer
}

// NewEntry - NewEntry function
func NewEntry(logger *logger, logLevel Level, args ...any) *LogEntry {
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

// Log - Log function prints the log message
func (le *LogEntry) Log(logLevel Level, args ...any) {
	le.log(logLevel, args...)
}

// LogF - LogF function prints the log message with the given format
func (le *LogEntry) LogF(logLevel Level, format string, args ...any) {
	le.log(logLevel, fmt.Sprintf(format, args...))
}

// log - log function is the utility function to print the log message in Log Function
func (le *LogEntry) log(logLevel Level, args ...any) {
	if le.logger.LogFilter.Filter(logLevel, args...) {
		// Add the args to the logMessage as a string
		le.logMessage = fmt.Sprint(args...)
		le.logLevel = logLevel
		// Add the log data
		le.logMessage = le.logger.LogFormatter.Format(le)

		fmt.Println(le.logMessage)
	}
}

// Info - Info function prints the log message with INFO level
func (le *LogEntry) Info(args ...any) {
	le.Log(INF, args...)
}

// InfoF - InfoF function prints the log message with INFO level and the given format
func (le *LogEntry) InfoF(format string, args ...any) {
	le.LogF(INF, format, args...)
}

// Debug - Debug function prints the log message with DEBUG level
func (le *LogEntry) Debug(args ...any) {
	le.Log(DBG, args...)
}

// DebugF - DebugF function prints the log message with DEBUG level and the given format
func (le *LogEntry) DebugF(format string, args ...any) {
	le.LogF(DBG, format, args...)
}

// Warn - Warn function prints the log message with WARN level
func (le *LogEntry) Warn(args ...any) {
	le.Log(WRN, args...)
}

// WarnF - WarnF function prints the log message with WARN level and the given format
func (le *LogEntry) WarnF(format string, args ...any) {
	le.LogF(WRN, format, args...)
}

// Error - Error function prints the log message with ERROR level
func (le *LogEntry) Error(args ...any) {
	le.Log(ERR, args...)
}

// ErrorF - ErrorF function prints the log message with ERROR level and the given format
func (le *LogEntry) ErrorF(format string, args ...any) {
	le.LogF(ERR, format, args...)
}

// WithContext - WithContext function
func (le *LogEntry) WithContext(ctx context.Context) *LogEntry {
	le.context = ctx
	return le
}

// WithField - WithField function
func (le *LogEntry) WithField(key string, value any) *LogEntry {
	le.logData[key] = value
	return le
}

// WithFields - WithFields function
func (le *LogEntry) WithFields(data map[string]any) *LogEntry {
	for key, value := range data {
		le.logData[key] = value
	}

	return le
}

// WithTime - WithTime function
func (le *LogEntry) WithTime(t time.Time) *LogEntry {
	le.logTime = t
	return le
}

// With Identifier - WithIdentifier function
func (le *LogEntry) WithIdentifier(identifier string) *LogEntry {
	le.logger.Ident = identifier
	return le
}
