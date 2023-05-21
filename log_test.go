package nexlog

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("Level.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_Color(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Color(); got != tt.want {
				t.Errorf("Level.Color() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ident  string
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ident, tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_EnableCaller(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.EnableCaller()
		})
	}
}

func Test_logger_EnableDefaultFileLogHook(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.EnableDefaultFileLogHook()
		})
	}
}

func Test_logger_Log(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Log(tt.args.logLevel, tt.args.args...)
		})
	}
}

func Test_logger_LogF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.LogF(tt.args.logLevel, tt.args.format, tt.args.args...)
		})
	}
}

func Test_logger_Debug(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Debug(tt.args.args...)
		})
	}
}

func Test_logger_DebugF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.DebugF(tt.args.format, tt.args.args...)
		})
	}
}

func Test_logger_Info(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Info(tt.args.args...)
		})
	}
}

func Test_logger_InfoF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.InfoF(tt.args.format, tt.args.args...)
		})
	}
}

func Test_logger_Warn(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Warn(tt.args.args...)
		})
	}
}

func Test_logger_WarnF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.WarnF(tt.args.format, tt.args.args...)
		})
	}
}

func Test_logger_Error(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Error(tt.args.args...)
		})
	}
}

func Test_logger_ErrorF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.ErrorF(tt.args.format, tt.args.args...)
		})
	}
}

func Test_getCaller(t *testing.T) {
	tests := []struct {
		name string
		want *runtime.Frame
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCaller(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCaller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_SetLogLevel(t *testing.T) {
	type fields struct {
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
	type args struct {
		logLevel Level
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.SetLogLevel(tt.args.logLevel)
		})
	}
}

func Test_logger_SetLogFormatter(t *testing.T) {
	type fields struct {
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
	type args struct {
		logFormatter LogFormatter
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.SetLogFormatter(tt.args.logFormatter)
		})
	}
}

func Test_logger_SetLogFilter(t *testing.T) {
	type fields struct {
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
	type args struct {
		logFilter LogFilter
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.SetLogFilter(tt.args.logFilter)
		})
	}
}

func Test_logger_AddHook(t *testing.T) {
	type fields struct {
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
	type args struct {
		logHook LogHook
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.AddHook(tt.args.logHook)
		})
	}
}

func Test_logger_Fatal(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.Fatal(tt.args.args...)
		})
	}
}

func Test_logger_FatalF(t *testing.T) {
	type fields struct {
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
			l := &logger{
				LogLevel:                 tt.fields.LogLevel,
				LogFormatter:             tt.fields.LogFormatter,
				LogFilter:                tt.fields.LogFilter,
				LogHooks:                 tt.fields.LogHooks,
				Ident:                    tt.fields.Ident,
				EntryPool:                tt.fields.EntryPool,
				isCaller:                 tt.fields.isCaller,
				enableDefaultFileLogHook: tt.fields.enableDefaultFileLogHook,
				mu:                       tt.fields.mu,
			}
			l.FatalF(tt.args.format, tt.args.args...)
		})
	}
}
