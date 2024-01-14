package app

import (
	"context"
	"log/slog"
	"os"
	"reflect"
	"testing"

	"GoSolveTestTask/config"
)

func TestInitApp(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		{
			name: "nil config",
			args: args{
				cfg: nil,
			},
			want: &App{data: nil, Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))},
		},
		{
			name: "with config",
			args: args{
				cfg: &config.Config{LogLevel: slog.LevelError},
			},
			want: &App{data: nil, Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitApp(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_ReadInput(t *testing.T) {
	type fields struct {
		data   []int
		Logger *slog.Logger
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "1",
			fields: fields{
				data:   []int{},
				Logger: logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				data:   tt.fields.data,
				Logger: tt.fields.Logger,
			}
			a.ReadInput()
		})
	}
}

func TestApp_GetIndex(t *testing.T) {
	type fields struct {
		data   []int
		Logger *slog.Logger
	}
	type args struct {
		ctx    context.Context
		target int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "positive",
			fields: fields{
				data:   []int{100},
				Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
			},
			args: args{
				ctx:    nil,
				target: 100,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "negative_too_much",
			fields: fields{
				data:   []int{100},
				Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
			},
			args: args{
				ctx:    nil,
				target: -1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative_too_much",
			fields: fields{
				data:   []int{100},
				Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
			},
			args: args{
				ctx:    nil,
				target: 1_000_001,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				data:   tt.fields.data,
				Logger: tt.fields.Logger,
			}
			got, err := a.GetIndex(tt.args.ctx, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.GetIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App.GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIndex(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 100,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 10,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				nums:   []int{0, 100, 200},
				target: 111,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				nums:   []int{0, 1, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				nums:   []int{21},
				target: 20,
			},
			want: 0,
		},
		{
			name: "6",
			args: args{
				nums:   []int{0, 100, 200},
				target: 110,
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 15,
			},
			want: -1,
		},
		{
			name: "8",
			args: args{
				nums:   []int{0, 10, 20, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200},
				target: 1150,
			},
			want: 13,
		},
		{
			name: "9",
			args: args{
				nums:   []int{},
				target: 1150,
			},
			want: -1,
		},
		{
			name: "10",
			args: args{
				nums:   []int{10},
				target: 1150,
			},
			want: -1,
		},
		{
			name: "11",
			args: args{
				nums:   []int{100},
				target: 99,
			},
			want: 0,
		},
		{
			name: "12",
			args: args{
				nums:   []int{100},
				target: 105,
			},
			want: 0,
		},
		{
			name: "13",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 98,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndex(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
