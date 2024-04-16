package telemetry

import (
	"go.opentelemetry.io/otel/metric"
)

// ApiActiveClients returns a counter that is not monotonic.
func ApiActiveClients(m metric.Meter) metric.Int64UpDownCounter {
	int64SumNoMonotonic, err := m.Int64UpDownCounter("api.active.clients",
		metric.WithDescription("Api active clients"), metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64SumNoMonotonic
}

// ApiMissingPaymentErrorCount returns a counter of the missing payment errors
func ApiMissingPaymentErrorCount(m metric.Meter) metric.Int64Counter {
	int64SumNoMonotonic, err := m.Int64Counter("api.missing.payments.count",
		metric.WithDescription("Api missing payments counter"), metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64SumNoMonotonic
}
