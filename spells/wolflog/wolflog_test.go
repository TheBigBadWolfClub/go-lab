package wolflog

import (
	"strings"
	"testing"
)

type mockOutput struct {
	result string
}

func (p *mockOutput) Write(data []byte) (n int, err error) {
	p.result = string(data)
	return len(data), nil
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{
			name: "Level To String",
			l:    0,
			want: "unknown",
		}, {
			name: "Level To String",
			l:    INFO,
			want: "Info",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wolflog_Debug(t *testing.T) {
	defaultWriter := &mockOutput{}
	defaultLog := New(DEBUG)
	defaultLog.Output(defaultWriter)

	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		log    Wolflogger
		writer *mockOutput
		args   args
	}{
		{
			name:   "should log as DEBUG",
			log:    defaultLog,
			writer: defaultWriter,
			args: args{
				format: "%s",
				v:      []interface{}{"log message"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.log.Debug(tt.args.format, tt.args.v...)

			if !strings.Contains(tt.writer.result, "log message") {
				t.Errorf("logger did not write log message")
			}
		})
	}
}
