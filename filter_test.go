package nexlog

import (
	"reflect"
	"testing"
)

func Test_newDefaultNoFilter(t *testing.T) {
	tests := []struct {
		name string
		want LogFilter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultNoFilter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultNoFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultNoFilter_Filter(t *testing.T) {
	type args struct {
		logLevel Level
		args     []any
	}
	tests := []struct {
		name string
		dnf  *defaultNoFilter
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dnf := &defaultNoFilter{}
			if got := dnf.Filter(tt.args.logLevel, tt.args.args...); got != tt.want {
				t.Errorf("defaultNoFilter.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
