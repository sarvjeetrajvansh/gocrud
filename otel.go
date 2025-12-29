package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() func(context.Context) error {
	exporter, err := stdouttrace.New()
	if err != nil {
		panic(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(
			sdktrace.ParentBased(
				sdktrace.TraceIDRatioBased(0.1), // 10%
			),
		),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
