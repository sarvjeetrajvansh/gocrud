package observability

import (
	"context"
	"github.com/sarvjeetrajvansh/gocrud/internal/config"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer(cfg *config.Config) func(context.Context) error {
	// OTLP exporter → Jaeger
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(cfg.OtelEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	SamplingRatio, _ := strconv.ParseFloat(cfg.SamplingRatio, 64)
	tp := sdktrace.NewTracerProvider(
		// ✅ sample 10% of requests
		sdktrace.WithSampler(
			sdktrace.ParentBased(
				sdktrace.TraceIDRatioBased(SamplingRatio),
			),
		),
		sdktrace.WithBatcher(
			exporter,
			sdktrace.WithBatchTimeout(5*time.Second),
		),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(cfg.AppName),
			),
		),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
