package pkg

import (
	"context"
	"go.opentelemetry.io/otel/metric"
)

const (
	Histogram MetricType = "histogram"
)

type MetricID string

func (m MetricID) String() string {
	return string(m)
}

type MetricType string

type MetricOption interface {
	metric.InstrumentOption
	metric.HistogramOption
}

type MetricRecorder interface {
	RecordCtx(ctx context.Context, value int, opts ...MetricRecorderOption)
	Record(value int, opts ...MetricRecorderOption)
}

type MetricRecorderOption func()

type Metric interface {
	RecordCtx(ctx context.Context, mid MetricID, value int, attr ...MetricRecorderOption)
	Record(mid MetricID, value int, attr ...MetricRecorderOption)
}

type metrics struct {
	meter         metric.Meter
	activeMetrics map[MetricID]MetricRecorder
}

func NewMetrics(meter metric.Meter) Metric {
	return &metrics{
		meter:         meter,
		activeMetrics: make(map[MetricID]MetricRecorder),
	}
}

func (m metrics) Register(t MetricType, name MetricID, opts ...MetricOption) error {

	switch t {
	case Histogram:
		m.activeMetrics[name], _ = newHistogram(m.meter, name, opts...)
	}

	return nil
}

func (m metrics) RecordCtx(ctx context.Context, mid MetricID, value int, attr ...MetricRecorderOption) {
	m.activeMetrics[mid].RecordCtx(ctx, value, attr...)
}

func (m metrics) Record(mid MetricID, value int, attr ...MetricRecorderOption) {
	m.activeMetrics[mid].Record(value, attr...)
}
