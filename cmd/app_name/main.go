package main

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"template/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg)

}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	if cfg.Logger.Level == nil {
		var level slog.Level
		if cfg.Env != "prod" {
			level = slog.LevelDebug.Level()
		} else {
			level = slog.LevelInfo.Level()
		}
		cfg.Logger.Level = &level
	}

	switch cfg.Env {
	case "local":
		log = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource: cfg.Logger.ShowPathCall,
				Level:     cfg.Logger.Level,
			}),
		)
	case "dev":
		log = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource: cfg.Logger.ShowPathCall,
				Level:     cfg.Logger.Level,
			}),
		)
	case "prod":
		log = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource: cfg.Logger.ShowPathCall,
				Level:     cfg.Logger.Level,
			}),
		)
	}

	return log
}
