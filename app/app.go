package app

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"GoSolveTestTask/config"
)

const allowedDeviationPercent = 10

type App struct {
	data   []int
	Logger *slog.Logger
}

func InitApp(cfg *config.Config) *App {
	logLevel := slog.LevelDebug
	if cfg != nil {
		logLevel = cfg.LogLevel.Level()
	}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return &App{Logger: logger}
}

func (a *App) ReadInput() error {
	timeStart := time.Now()

	result := make([]int, 0, 1_000_010)
	var integer int
	for {
		if _, err := fmt.Scanln(&integer); err != nil {
			a.Logger.Error("ReadInput", "err", err)
			if err != io.EOF {
				return err
			}
			break
		}

		result = append(result, integer)
	}
	a.Logger.Info("read records from input", "len", len(result), "in", time.Since(timeStart).String())

	a.data = result

	return nil
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
		if abs(target-nums[0]) <= target/allowedDeviationPercent {
			return 0
		}

		return -1
	}

	if abs(target-nums[med]) <= target/allowedDeviationPercent {
		return med
	} else if next := med + 1; next < len(nums) && nums[next]-target <= target/allowedDeviationPercent {
		return next
	}

	return -1
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
