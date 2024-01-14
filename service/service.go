package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Server struct {
	Logger      *slog.Logger
	DataService DataService
}

type DataService interface {
	GetIndex(ctx context.Context, target int) (int, error)
}

func New(log *slog.Logger, ds DataService) *Server {
	return &Server{Logger: log, DataService: ds}
}

func (s *Server) FindIndex(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Index        int    `json:"index"`
		ErrorMessage string `json:"error_message,omitempty"`
	}

	target, err := strconv.Atoi(chi.URLParam(r, "value"))
	if err != nil {
		resp := response{ErrorMessage: fmt.Sprintf("target number %v is not an integer, err: %v", target, err)}
		s.renderResponse(w, resp)
		return
	}
	s.Logger.Info("got request to find target", "integer", target)

	res, err := s.DataService.GetIndex(r.Context(), target)
	if err != nil {
		resp := response{ErrorMessage: err.Error()}
		s.renderResponse(w, resp)
		return
	}

	if res == -1 {
		resp := response{ErrorMessage: fmt.Sprintf("target number %v is not found", target)}
		s.renderResponse(w, resp)
		return
	}
	resp := response{Index: res}
	s.renderResponse(w, resp)
}

func (s *Server) renderResponse(w http.ResponseWriter, data any) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		s.Logger.Error("renderResponse", "err", err)
	}
	s.Logger.Debug("renderResponse", "resp", data)
}
