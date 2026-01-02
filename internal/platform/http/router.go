package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/sarvjeetrajvansh/gocrud/internal/observability"
	"github.com/sarvjeetrajvansh/gocrud/internal/user"
)

type Router struct {
	handler http.Handler
}

// NewRouter builds and returns the HTTP handler
func NewRouter(
	appName string,
	logger *slog.Logger,
	userHandler *user.Handler,
) http.Handler {

	r := chi.NewRouter()

	// ---- Core middleware ----
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// ---- Logging ----
	r.Use(middleware.RequestLogger(&observability.SlogFormatter{
		Logger: logger,
	}))

	// ---- Routes ----
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUserByID)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	// ---- OpenTelemetry wrapper ----
	return otelhttp.NewHandler(r, appName)
}
