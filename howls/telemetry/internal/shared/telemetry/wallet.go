package telemetry

import "go.opentelemetry.io/otel/metric"

func WalletBalanceObservable(m metric.Meter, callback metric.Int64Callback) metric.Int64ObservableGauge {
	int64ObservableGauge, err := m.Int64ObservableGauge("wallet.balance",
		metric.WithDescription("Api stuff"),
		metric.WithInt64Callback(callback),
		metric.WithUnit("un"))
	if err != nil {
		panic(err)
	}

	return int64ObservableGauge
}
