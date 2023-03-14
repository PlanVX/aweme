package otel

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"time"
)

// TracerProvider returns a new open telemetry tracer provider
func TracerProvider(conf *config.Config, lf fx.Lifecycle) (*trace.TracerProvider, error) {

	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(conf.Uptrace.DSN),
		uptrace.WithServiceName(conf.Uptrace.Service),
		uptrace.WithServiceVersion(conf.Uptrace.Version),
		uptrace.WithDeploymentEnvironment(conf.Uptrace.Environment),
	)
	tp := uptrace.TracerProvider()
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
