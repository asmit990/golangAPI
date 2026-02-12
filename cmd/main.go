package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/asmit990/golangAPI/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user=postgres password=postgres dbname=app sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx)

	// database
	logger.Info("connected to database", "dsn", cfg.db.dsn)
	api := application{
		config: cfg,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		slog.Error("server is failed", "error", err)
		os.Exit(1)
	}
}
