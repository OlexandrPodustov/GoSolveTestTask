package service

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
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

type mockDataService struct{}

func (m *mockDataService) GetIndex(ctx context.Context, target int) (int, error) { return 1, nil }

func TestServer_FindIndex(t *testing.T) {
	type fields struct {
		Logger      *slog.Logger
		DataService DataService
	}
	type args struct {
		r *http.Request
	}

	route := "/endpoint"
	var mds mockDataService

	tests := []struct {
		name     string
		fields   fields
		argsFn   func(t *testing.T) args
		wantCode int
	}{
		{
			name: "found",
			fields: fields{
				Logger:      slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
				DataService: &mds,
			},
			argsFn: func(t *testing.T) args {
				t.Helper()

				req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, route, new(bytes.Buffer))
				require.NoError(t, err)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("value", "111")
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

				return args{r: req}
			},
			wantCode: http.StatusOK,
		},
		{
			name: "bad_param",
			fields: fields{
				Logger:      slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
				DataService: &mds,
			},
			argsFn: func(t *testing.T) args {
				t.Helper()

				req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, route, new(bytes.Buffer))
				require.NoError(t, err)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("value", "aa")
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

				return args{r: req}
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:      tt.fields.Logger,
				DataService: tt.fields.DataService,
			}

			w := httptest.NewRecorder()
			arg := tt.argsFn(t)
			s.FindIndex(w, arg.r)

			require.Equal(t, tt.wantCode, w.Code)
		})
	}
}
