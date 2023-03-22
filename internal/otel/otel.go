package otel

import (
	"context"
	"time"

	"github.com/PlanVX/aweme/internal/config"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.18.0"
)

// TracerProvider constructs a new trace provider.
func TracerProvider(conf *config.Config) (*trace.TracerProvider, error) {
	if conf.Otel.Enabled == false {
		return trace.NewTracerProvider(), nil
	}
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(conf.Otel.Endpoint),
		otlptracegrpc.WithCompressor("gzip"),
	)
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.Otel.Service),
			semconv.DeploymentEnvironment(conf.Otel.Environment),
		),
		),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func stopTracerProvider(ctx context.Context, tp *trace.TracerProvider) error {
	return tp.Shutdown(ctx)
}

// MeterProvider constructs a new meter provider.
func MeterProvider(conf *config.Config) (*metric.MeterProvider, error) {
	if conf.Otel.Enabled == false {
		return metric.NewMeterProvider(), nil
	}
	options := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(conf.Otel.Endpoint),
		otlpmetricgrpc.WithCompressor("gzip"),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithTemporalitySelector(preferDeltaTemporalitySelector),
	}
	exporter, err := otlpmetricgrpc.New(context.Background(), options...)
	if err != nil {
		return nil, err
	}
	reader := metric.NewPeriodicReader(
		exporter,
		metric.WithInterval(15*time.Second),
	)
	provider := metric.NewMeterProvider(
		metric.WithReader(reader),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.Otel.Service),
			semconv.DeploymentEnvironment(conf.Otel.Environment),
		)))
	global.SetMeterProvider(provider)
	if err := runtime.Start(); err != nil {
		return nil, err
	}
	return provider, nil
}

func stopMeterProvider(ctx context.Context, tp *metric.MeterProvider) error {
	return tp.Shutdown(ctx)
}

func preferDeltaTemporalitySelector(kind metric.InstrumentKind) metricdata.Temporality {
	switch kind {
	case metric.InstrumentKindCounter,
		metric.InstrumentKindObservableCounter,
		metric.InstrumentKindHistogram:
		return metricdata.DeltaTemporality
	default:
		return metricdata.CumulativeTemporality
	}
}
