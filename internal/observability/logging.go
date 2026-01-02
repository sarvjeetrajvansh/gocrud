package observability

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel/trace"
)

// SlogFormatter plugs into chi RequestLogger
type SlogFormatter struct {
	Logger *slog.Logger
}

// NewLogEntry is called at request start
func (f *SlogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return &SlogEntry{
		logger: f.Logger,
		req:    r,
		start:  time.Now(),
	}
}

// SlogEntry logs after response is written
type SlogEntry struct {
	logger *slog.Logger
	req    *http.Request
	start  time.Time
}

// Write runs after handler completes
func (e *SlogEntry) Write(
	status, bytes int,
	header http.Header,
	elapsed time.Duration,
	extra interface{},
) {
	// trace_id from OpenTelemetry
	span := trace.SpanFromContext(e.req.Context())
	traceID := span.SpanContext().TraceID().String()
	if !span.SpanContext().HasTraceID() {
		traceID = "none"
	}

	// request_id from chi
	requestID := middleware.GetReqID(e.req.Context())
	if requestID == "" {
		requestID = "none"
	}

	e.logger.Info(
		"http request",
		"method", e.req.Method,
		"path", e.req.URL.Path,
		"status", status,
		"duration", elapsed,
		"request_id", requestID,
		"trace_id", traceID,
	)
}

// Panic logs panics
func (e *SlogEntry) Panic(v interface{}, stack []byte) {
	e.logger.Error(
		"panic",
		"panic", v,
	)
}
