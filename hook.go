package nexlog

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// LogHook is an interface that encapsulates actions to be taken after a log entry is created.
type LogHook interface {
	Fire(le *LogEntry) error
}

// defaultJsonFileLogHook is an implementation of LogHook that writes log entries to a JSON file.
type defaultJsonFileLogHook struct{ mu sync.Mutex }

// newDefaultJsonFileLogHook creates and returns a new instance of defaultJsonFileLogHook.
func newDefaultJsonFileLogHook() LogHook {
	return &defaultJsonFileLogHook{}
}

// Fire method for defaultJsonFileLogHook writes the provided log entry to a file named "app.log" in JSON format.
// If the file does not exist, it creates a new one. If it exists, it appends the new log entry to the file.
func (djf *defaultJsonFileLogHook) Fire(le *LogEntry) error {
	djf.mu.Lock()
	defer djf.mu.Unlock()
	// Create a new file called app.log if it doesn't exist
	if _, err := os.Stat("app.log"); os.IsNotExist(err) {
		file, err := os.Create("app.log")
		if err != nil {
			return err
		}
		defer file.Close()
	}

	// Open the file in append mode
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	JsonEntry := map[string]any{
		"level": le.logLevel.String(),
		"time":  le.logTime,
		"msg":   le.rowMessage,
	}

	if le.caller != nil {
		JsonEntry["path"] = fmt.Sprintf("%s:%d", le.caller.File, le.caller.Line)
	}

	jsonData, err := json.Marshal(JsonEntry)
	if err != nil {
		return err
	}

	logMsg := string(jsonData) + "\n"
	// Write the log message to the file
	if _, err := file.WriteString(logMsg); err != nil {
		return err
	}

	return nil
}
