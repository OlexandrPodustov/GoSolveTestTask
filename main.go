package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"GoSolveTestTask/config"
	"GoSolveTestTask/service"
)

func main() {
	configFile := flag.String("config", "./config.json", "Configuration file in JSON-format")
	flag.Parse()

	cfg, err := config.Load(configFile)
	if err != nil {
		log.Default().Println(err)

		cfg = &config.Config{
			Port:     8080,
			LogLevel: slog.LevelDebug,
		}
	}

	app, err := service.InitApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	app.ReadInput()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/endpoint/{value}", app.FindIndex)

	if err := http.ListenAndServe(fmt.Sprint(":", cfg.Port), router); err != nil {
		log.Fatal(err)
	}
}
