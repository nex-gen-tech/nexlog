package nexlog

import (
	"strconv"
)

type LogFormatter interface {
	Format(le *LogEntry) string
}

type defaultTextFormatter struct{}

func NewDefaultTextFormatter() LogFormatter {
	return &defaultTextFormatter{}
}

func (dtf *defaultTextFormatter) Format(le *LogEntry) string {
	line := le.caller.File + ":" + strconv.Itoa(le.caller.Line) + " " + le.logLevel.String() + " " + le.timeStamp.Format("2006-01-02 15:04:05") + " " + le.logMessage

	return line
}
