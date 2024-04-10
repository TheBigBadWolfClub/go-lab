package presentation

import (
	"context"
	"github.com/TheBigBadWolfClub/go-lab/howls/instrumentation/internal/core"
	"github.com/TheBigBadWolfClub/go-lab/howls/instrumentation/internal/shared/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"net/http"
	"strconv"
	"time"
)

type API interface {
	DiceRoller(http.ResponseWriter, *http.Request)
}

type api struct {
	diceRollerSrv             core.DiceRoller
	telemetry                 telemetry.Telemetry
	rollCount                 metric.Int64Counter
	currentRollValue          metric.Int64UpDownCounter
	rollValueDistribution     metric.Int64Histogram
	diceSrvResponseTimeMetric metric.Int64Histogram
	diceSrcCalls              metric.Int64Counter
}

func NewAPI(tel telemetry.Telemetry) API {

	return &api{
		telemetry:     tel,
		diceRollerSrv: core.NewDiceRoller(tel.Meter()),

		// dice metrics
		rollCount:             telemetry.DoubleDiceRollCount(tel.Meter()),
		currentRollValue:      telemetry.DoubleDiceValue(tel.Meter()),
		rollValueDistribution: telemetry.DoubleDiceValueDistribution(tel.Meter()),

		// dice service metrics
		diceSrvResponseTimeMetric: telemetry.DiceServiceResponseTime(tel.Meter()),
		diceSrcCalls:              telemetry.DiceServiceTotalCalls(tel.Meter()),
	}
}

func (a *api) DiceRoller(w http.ResponseWriter, r *http.Request) {

	spanCtx, span := a.telemetry.Tracer().Start(r.Context(), "roll")
	defer span.End()

	nDices := 2
	ints := make(chan int, nDices)
	for i := 0; i < nDices; i++ {
		go func() {
			ints <- a.timedSrvCall(spanCtx)
		}()
	}

	totalValue := 0
	for i := 0; i < nDices; i++ {
		select {
		case roll := <-ints:
			totalValue += roll
		}
	}

	rollValueAttr := attribute.Int("roll.value", totalValue)
	span.SetAttributes(rollValueAttr)

	a.metricCounter(spanCtx, rollValueAttr)
	a.metricGauge(spanCtx, totalValue)
	a.metricHistogram(spanCtx, totalValue)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strconv.Itoa(totalValue)))
}

func (a *api) timedSrvCall(ctx context.Context) int {
	start := time.Now()

	roll := a.diceRollerSrv.Roll(ctx)

	elapsed := time.Since(start)
	a.diceSrvResponseTimeMetric.Record(ctx, elapsed.Milliseconds())
	a.diceSrcCalls.Add(ctx, 1)

	return roll
}

func (a *api) metricCounter(ctx context.Context, rollValueAttr attribute.KeyValue) {
	a.rollCount.Add(ctx, 1, metric.WithAttributes(rollValueAttr))
}

func (a *api) metricGauge(ctx context.Context, roll int) {
	a.currentRollValue.Add(ctx, int64(roll))
}

func (a *api) metricHistogram(ctx context.Context, roll int) {
	a.rollValueDistribution.Record(ctx, int64(roll))
}
