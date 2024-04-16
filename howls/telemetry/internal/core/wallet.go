package core

import (
	"context"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/shared/telemetry"
	"go.opentelemetry.io/otel/metric"
	"sync"
)

var mu sync.Mutex

type Wallet interface {
	Deposit()
	Withdraw()
	Balance() int
}

type wallet struct {
	balance int
}

func NewWallet(m metric.Meter) Wallet {
	w := &wallet{
		balance: 10,
	}
	telemetry.WalletBalanceObservable(m, w.BalanceMetric)
	return w
}

func (w *wallet) Deposit() {
	mu.Lock()
	defer mu.Unlock()
	w.balance = 10
}

func (w *wallet) Withdraw() {
	mu.Lock()
	defer mu.Unlock()
	w.balance -= 1
}

func (w *wallet) Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return w.balance
}

func (w *wallet) BalanceMetric(_ context.Context, valueObserver metric.Int64Observer) error {
	mu.Lock()
	defer mu.Unlock()
	valueObserver.Observe(int64(w.balance))
	return nil
}
