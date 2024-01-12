package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	"GoSolveTestTask/config"
)

type App struct {
	data   []int
	Logger *slog.Logger
}

type Response struct {
	Index        int    `json:"index"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func InitApp(cfg *config.Config) (*App, error) {
	opts := &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return &App{Logger: logger}, nil
}

func (a *App) ReadInput() {
	var integer int

	timeStart := time.Now()

	result := make([]int, 0, 1_000_010)
	for {
		_, err := fmt.Scanln(&integer)
		if err != nil {
			a.Logger.Info("ReadInput", "err", err)
			break
		}

		result = append(result, integer)
	}
	a.Logger.Info("read records from input", "len", len(result), "in", time.Since(timeStart).String())

	a.data = result
}

func (a *App) FindIndex(w http.ResponseWriter, r *http.Request) {
	target, err := strconv.Atoi(chi.URLParam(r, "value"))
	if err != nil {
		resp := Response{ErrorMessage: fmt.Sprintf("target number %v is not an integer, err: %v", target, err)}
		a.renderResponse(w, resp)
		return
	}
	a.Logger.Info("target", "integer", target)

	if target < 0 || target > 1000000 { // min and max check - and return fast if impossible
		resp := Response{ErrorMessage: fmt.Sprintf("target number %v is out of bounds", target)}
		a.renderResponse(w, resp)
		return
	}

	res := a.findIndex(target)
	if res == -1 {
		resp := Response{ErrorMessage: fmt.Sprintf("target number %v is not found", target)}
		a.renderResponse(w, resp)
		return
	}
	resp := Response{Index: res}
	a.renderResponse(w, resp)
}

func (a *App) renderResponse(w http.ResponseWriter, data Response) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		a.Logger.Error("renderResponse", "err", err)
	}
	a.Logger.Debug("renderResponse", "resp", data)
}

func (a *App) findIndex(target int) int {
	return -1
}
