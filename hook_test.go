package nexlog

import (
	"reflect"
	"sync"
	"testing"
)

func Test_newDefaultJsonFileLogHook(t *testing.T) {
	tests := []struct {
		name string
		want LogHook
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultJsonFileLogHook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultJsonFileLogHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultJsonFileLogHook_Fire(t *testing.T) {
	type fields struct {
		mu sync.Mutex
	}
	type args struct {
		le *LogEntry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			djf := &defaultJsonFileLogHook{
				mu: tt.fields.mu,
			}
			if err := djf.Fire(tt.args.le); (err != nil) != tt.wantErr {
				t.Errorf("defaultJsonFileLogHook.Fire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
