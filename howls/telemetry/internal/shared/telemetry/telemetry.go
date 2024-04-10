package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Telemetry interface {
	Tracer() trace.Tracer
	Meter() metric.Meter
}

type telemetry struct {
	tracer trace.Tracer
	meter  metric.Meter
}

func NewTelemetry(appID string) Telemetry {
	return &telemetry{
		tracer: otel.Tracer(appID),
		meter:  otel.Meter(appID),
	}
}

func (t *telemetry) Tracer() trace.Tracer {
	return t.tracer
}

func (t *telemetry) Meter() metric.Meter {
	return t.meter
}
