package nexlog

import (
	"bytes"
	"context"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func Test_newEntry(t *testing.T) {
	type args struct {
		logger   *logger
		logLevel Level
		args     []any
	}
	tests := []struct {
		name string
		args args
		want *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEntry(tt.args.logger, tt.args.logLevel, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogEntry_Log(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		logLevel Level
		args     []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Log(tt.args.logLevel, tt.args.args...)
		})
	}
}

func TestLogEntry_LogF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		logLevel Level
		format   string
		args     []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.LogF(tt.args.logLevel, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_log(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		logLevel Level
		args     []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.log(tt.args.logLevel, tt.args.args...)
		})
	}
}

func TestLogEntry_Info(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Info(tt.args.args...)
		})
	}
}

func TestLogEntry_InfoF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.InfoF(tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_Debug(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Debug(tt.args.args...)
		})
	}
}

func TestLogEntry_DebugF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.DebugF(tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_Warn(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Warn(tt.args.args...)
		})
	}
}

func TestLogEntry_WarnF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.WarnF(tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_Error(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Error(tt.args.args...)
		})
	}
}

func TestLogEntry_ErrorF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.ErrorF(tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_Fatal(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.Fatal(tt.args.args...)
		})
	}
}

func TestLogEntry_FatalF(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			le.FatalF(tt.args.format, tt.args.args...)
		})
	}
}

func TestLogEntry_WithContext(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			if got := le.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogEntry.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogEntry_WithField(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			if got := le.WithField(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogEntry.WithField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogEntry_WithFields(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		data map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			if got := le.WithFields(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogEntry.WithFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogEntry_WithTime(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			if got := le.WithTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogEntry.WithTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogEntry_WithIdentifier(t *testing.T) {
	type fields struct {
		logger     *logger
		logLevel   Level
		logTime    time.Time
		rowMessage string
		logMessage string
		logData    map[string]any
		caller     *runtime.Frame
		timeStamp  time.Time
		context    context.Context
		buffer     *bytes.Buffer
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *LogEntry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			le := &LogEntry{
				logger:     tt.fields.logger,
				logLevel:   tt.fields.logLevel,
				logTime:    tt.fields.logTime,
				rowMessage: tt.fields.rowMessage,
				logMessage: tt.fields.logMessage,
				logData:    tt.fields.logData,
				caller:     tt.fields.caller,
				timeStamp:  tt.fields.timeStamp,
				context:    tt.fields.context,
				buffer:     tt.fields.buffer,
			}
			if got := le.WithIdentifier(tt.args.identifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogEntry.WithIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
