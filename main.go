package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sarvjeetrajvansh/gocrud/handlers"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"github.com/sarvjeetrajvansh/gocrud/storage"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"log/slog"
	"net/http"
	"os"
)

const PORT string = ":8080"

func main() {

	shutdown := initTracer()
	defer shutdown(context.Background())

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	userStore := storage.NewUserstore()
	userService := service.NewUserservice(userStore)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&SlogFormatter{
		logger: logger,
	}))
	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsers(userService))
		r.Post("/", handlers.CreateUser(userService))
		r.Get("/{id}", handlers.GetUserByID(userService))
		r.Put("/{id}", handlers.UpdateUser(userService))
		r.Delete("/{id}", handlers.DeleteUser(userService))
	})

	handler := otelhttp.NewHandler(r, "gocrud")

	http.ListenAndServe(PORT, handler)

}
