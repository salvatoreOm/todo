package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"todo/internal/database"
	"todo/internal/handler"
	"todo/internal/repository"
	"todo/internal/service"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// run wires the layers together: database -> repository -> service -> handler.
func run() error {
	// Load .env if it exists; real environment variables still win.
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is not set (see .env.example)")
	}

	db, err := database.Connect(dsn)
	if err != nil {
		return err
	}
	if sqlDB, err := db.DB(); err == nil {
		defer sqlDB.Close()
	}

	repo := repository.NewTodoRepository(db)
	svc := service.NewTodoService(repo)

	// Ctrl+C cancels the context, which stops any running DB query.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	return handler.NewRootCmd(svc).ExecuteContext(ctx)
}
