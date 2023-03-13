package otel

import "go.uber.org/fx"

// Module is a Fx module that initializes the OpenTelemetry MeterProvider and TracerProvider.
var Module = fx.Module("otel",
	fx.Invoke(
		fx.Annotate(
			MeterProvider,
			fx.OnStop(stopMeterProvider),
		),
		fx.Annotate(
			TracerProvider,
			fx.OnStop(stopTracerProvider),
		),
	),
)
