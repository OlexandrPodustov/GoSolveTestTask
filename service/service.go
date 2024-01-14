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

	res := findIndex(a.data, target)
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

func findIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	med := len(nums) / 2

	for low <= high {
		if nums[med] == target {
			return med
		} else if target > nums[med] {
			low = med + 1
		} else {
			high = med - 1
		}
		med = (high + low) / 2
	}

	return withinTen(nums, target, med)
}

func withinTen(nums []int, target int, med int) int {
	if len(nums) == 1 {
		if target-nums[0] <= target/10 {
			return 0
		} else if nums[0]-target <= target/10 {
			return 0
		}

		return -1
	}

	if prev := med - 1; prev >= 0 && target-nums[prev] <= target/10 {
		return prev
	} else if next := med + 1; next < len(nums) && nums[next]-target <= target/10 {
		return next
	}

	return -1
}
