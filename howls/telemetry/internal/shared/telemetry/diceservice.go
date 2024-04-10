package telemetry

import "go.opentelemetry.io/otel/metric"

func DiceServiceTotalCalls(m metric.Meter) metric.Int64Counter {
	int64Counter, err := m.Int64Counter("dice.srv.calls",
		metric.WithDescription("The count of number of dice service calls"),
		metric.WithUnit("calls"))
	if err != nil {
		panic(err)
	}

	return int64Counter
}

func DiceServiceResponseTime(m metric.Meter) metric.Int64Histogram {
	int64Histogram, err := m.Int64Histogram("dice.srv.response.time",
		metric.WithDescription("The distribution of 2 dice rolls"),
		metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64Histogram
}
