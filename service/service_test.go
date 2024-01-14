package service

import (
	"log/slog"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		log *slog.Logger
		ds  DataService
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "1",
			args: args{
				log: &slog.Logger{},
				ds:  nil,
			},
			want: &Server{
				Logger:      &slog.Logger{},
				DataService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.log, tt.args.ds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
