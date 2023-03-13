package otel

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/fx"
	"time"
)

func TracerProvider(conf *config.Config, lf fx.Lifecycle) (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(conf.Otel.Endpoint),
	)
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.Otel.Service),
			attribute.String("environment", conf.Otel.Environment),
		),
		),
	)
	lf.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				// Cleanly shutdown and flush telemetry when the application exits.
				// Do not make the application hang when it is shutdown.
				ctx, cancel := context.WithTimeout(ctx, time.Second*5)
				defer cancel()
				return tp.Shutdown(ctx)
			},
		})
	otel.SetTracerProvider(tp)
	return tp, nil
}
