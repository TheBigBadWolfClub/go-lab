package telemetry

import "go.opentelemetry.io/otel/metric"

// DoubleDiceRollCount returns a counter that is monotonic. used as Counter
func DoubleDiceRollCount(m metric.Meter) metric.Int64Counter {
	int64Counter, err := m.Int64Counter("double.dice.roll.count",
		metric.WithDescription("The count of number of dice rolls"),
		metric.WithUnit("rolls"))
	if err != nil {
		panic(err)
	}

	return int64Counter
}

// DoubleDiceValueDistribution returns a histogram with distribution of 2 dice rolls
func DoubleDiceValueDistribution(m metric.Meter) metric.Int64Histogram {
	int64Histogram, err := m.Int64Histogram("double.dice.roll.value.distribution",
		metric.WithExplicitBucketBoundaries([]float64{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}...),
		metric.WithDescription("The distribution of 2 dice rolls"), metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64Histogram
}
