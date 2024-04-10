package core

import (
	"context"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/shared/telemetry"
	"go.opentelemetry.io/otel/metric"
	"math/rand"
	"time"
)

type DiceRoller interface {
	Roll(ctx context.Context) int
}

type dice struct {
	valueDistribution metric.Int64Histogram
}

func (d dice) Roll(ctx context.Context) int {

	sleepy := time.Duration(rand.Intn(1000))
	time.Sleep(sleepy * time.Millisecond)

	rollValue := rand.Intn(5) + 1

	d.valueDistribution.Record(ctx, int64(rollValue))
	return rollValue
}

func NewDiceRoller(m metric.Meter) DiceRoller {
	return &dice{
		valueDistribution: telemetry.SingleDiceValueDistribution(m),
	}
}
