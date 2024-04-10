package pkg

import (
	"context"
	"go.opentelemetry.io/otel/metric"
)

func WithDescription(desc string) MetricOption {
	return metric.WithDescription(desc)
}

func WithUnit(unit string) MetricOption {
	return metric.WithUnit(unit)
}

func WithExplicitBucketBoundaries(buckets ...float64) metric.HistogramOption {

	return metric.WithExplicitBucketBoundaries(buckets...)
}

func newHistogram(m metric.Meter, name MetricID, opts ...MetricOption) (MetricRecorder, error) {

	int64Histogram, err := m.Int64Histogram(name.String(),
		metric.WithExplicitBucketBoundaries([]float64{1, 2, 3, 4, 5, 6}...),
		metric.WithDescription("The distribution of dice rolls"),
		metric.WithUnit("un"))
	if err != nil {
		return nil, err
	}

	return &histogram{
		histogram: int64Histogram,
	}, nil
}

type histogram struct {
	histogram metric.Int64Histogram
	_type     interface{}
}

func (h histogram) RecordCtx(ctx context.Context, value int, attr ...MetricRecorderOption) {
	h.histogram.Record(ctx, int64(value))
}

func (h histogram) Record(value int, attr ...MetricRecorderOption) {
	h.histogram.Record(context.Background(), int64(value))
}
