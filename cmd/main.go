package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("server is failed", "error", err)
		os.Exit(1)
	}
}
