package app

import (
	"GoSolveTestTask/config"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type App struct {
	data   []int
	Logger *slog.Logger
}

func InitApp(cfg *config.Config) (*App, error) {
	opts := &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return &App{Logger: logger}, nil
}

func (a *App) ReadInput() {
	timeStart := time.Now()

	result := make([]int, 0, 1_000_010)
	var integer int
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

func (a *App) GetIndex(ctx context.Context, target int) (int, error) {
	if target < 0 || target > 1_000_000 { // min and max check - and return fast if impossible
		return 0, fmt.Errorf("target number %v is out of bounds", target)
	}
	return findIndex(a.data, target), nil
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

	if target-nums[med] <= target/10 {
		return med
	} else if prev := med - 1; prev >= 0 && target-nums[prev] <= target/10 {
		return prev
	} else if next := med + 1; next < len(nums) && nums[next]-target <= target/10 {
		return next
	}

	return -1
}
