package nexlog

import (
	"reflect"
	"testing"
)

func Test_newDefaultTextFormatter(t *testing.T) {
	tests := []struct {
		name string
		want LogFormatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultTextFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultTextFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultTextFormatter_Format(t *testing.T) {
	type args struct {
		le *LogEntry
	}
	tests := []struct {
		name string
		dtf  *defaultTextFormatter
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtf := &defaultTextFormatter{}
			if got := dtf.Format(tt.args.le); got != tt.want {
				t.Errorf("defaultTextFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
