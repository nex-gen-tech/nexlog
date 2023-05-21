package nexlog

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

// LogFormatter is an interface that provides a mechanism to format log entries into strings.
type LogFormatter interface {
	Format(le *LogEntry) string
}

// defaultTextFormatter is an implementation of LogFormatter interface that formats log entries in a default textual format.
type defaultTextFormatter struct{}

// newDefaultTextFormatter creates and returns a new instance of defaultTextFormatter.
func newDefaultTextFormatter() LogFormatter {
	return &defaultTextFormatter{}
}

// Format method for defaultTextFormatter formats the given log entry into a string.
// The formatted string includes log identifier, colored log level, log time, log message, and the location in the source file of the log call.
func (dtf *defaultTextFormatter) Format(le *LogEntry) string {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	logLvl := le.logLevel.Color()

	logMsg := fmt.Sprintf(
		"[%s] %s %s | %s",
		le.logger.Ident,
		logLvl,
		logTime,
		le.logMessage,
	)

	if le.caller != nil {
		fileLoc := fmt.Sprintf("%s:%d", le.caller.File, le.caller.Line)
		fileLoc = color.FgMagenta.Render(fileLoc)
		logMsg += fmt.Sprintf(" %s", fileLoc)
	}

	return logMsg
}
