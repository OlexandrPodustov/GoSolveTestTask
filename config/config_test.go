package config

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	type args struct {
		filename *string
	}

	fileAbsent := "test.json"
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name:    "no_config",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "config_file_absent",
			args: args{
				filename: &fileAbsent,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "bad_config_file_present",
			args: args{
				filename: func() *string {
					filePresent := "present.json"

					f, err := os.CreateTemp("", filePresent)
					if err != nil {
						t.Errorf("TempFile(..., %q) error: %v", filePresent, err)
						return nil
					}
					f.Write([]byte(`}`))

					fn := f.Name()
					return &fn
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "config_file_present_empty",
			args: args{
				filename: func() *string {
					filePresent := "present.json"

					f, err := os.CreateTemp("", filePresent)
					if err != nil {
						t.Errorf("TempFile(..., %q) error: %v", filePresent, err)
						return nil
					}
					f.Write([]byte(`{}`))

					fn := f.Name()
					return &fn
				}(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "config_file_present",
			args: args{
				filename: func() *string {
					filePresent := "present.json"

					f, err := os.CreateTemp("", filePresent)
					if err != nil {
						t.Errorf("TempFile(..., %q) error: %v", filePresent, err)
						return nil
					}
					f.Write([]byte(`{"log_level":"debug"}`))

					fn := f.Name()
					return &fn
				}(),
			},
			want:    &Config{LogLevel: slog.LevelDebug},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.filename != nil {
				defer os.Remove(*tt.args.filename)
			}

			got, err := Load(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
