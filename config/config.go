package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Port     int
	LogLevel slog.Leveler
}

func Load(filename *string) (*Config, error) {
	if filename == nil {
		return nil, fmt.Errorf("no config file")
	}
	fileContents, err := os.ReadFile(*filename)
	if err != nil {
		return nil, err
	}
	var config struct {
		Port     int    `json:"server_port"`
		LogLevel string `json:"log_level"`
	}
	if err := json.Unmarshal(fileContents, &config); err != nil {
		return nil, err
	}

	var lvl slog.LevelVar
	if err := lvl.UnmarshalText([]byte(config.LogLevel)); err != nil {
		return nil, err
	}

	return &Config{
		Port:     config.Port,
		LogLevel: lvl.Level(),
	}, nil
}
