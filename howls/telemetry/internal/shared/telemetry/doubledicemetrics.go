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

// DoubleDiceValue returns a counter that is not monotonic. used as Gauge
func DoubleDiceValue(m metric.Meter) metric.Int64UpDownCounter {

	int64SumNoMonotonic, err := m.Int64UpDownCounter("double.dice.roll.value",
		metric.WithDescription("The last value that was rolled"), metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64SumNoMonotonic
}

// SingleDiceValueDistribution returns a histogram with distribution of 1 dice rolls
func SingleDiceValueDistribution(m metric.Meter) metric.Int64Histogram {
	int64Histogram, err := m.Int64Histogram("single.dice.roll.value.distribution",
		metric.WithExplicitBucketBoundaries([]float64{1, 2, 3, 4, 5, 6}...),
		metric.WithDescription("The distribution of 1 dice rolls"), metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64Histogram
}
