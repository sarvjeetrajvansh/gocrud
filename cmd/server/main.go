package main

import (
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"log/slog"

	"github.com/sarvjeetrajvansh/gocrud/internal/config"
	"github.com/sarvjeetrajvansh/gocrud/internal/observability"
	router "github.com/sarvjeetrajvansh/gocrud/internal/platform/http"
	"github.com/sarvjeetrajvansh/gocrud/internal/storage/postgres"
	"github.com/sarvjeetrajvansh/gocrud/internal/user"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	// ---- Observability ----
	shutdown := observability.InitTracer(cfg)
	defer shutdown(context.Background())

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	// ---- Database ----
	db := postgres.NewGormDB(cfg.DBDSN)
	db.AutoMigrate(&user.User{})

	// ---- User wiring ----
	userRepo := postgres.NewUserRepo(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// ---- Router (THIS IS THE KEY FIX) ----
	httpHandler := router.NewRouter(
		cfg.AppName,
		logger,
		userHandler,
	)

	// ---- Server ----
	addr := ":" + cfg.HTTPPort
	http.ListenAndServe(addr, httpHandler)
}
