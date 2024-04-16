package presentation

import (
	"context"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/core"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/shared/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API interface {
	DiceRoller(http.ResponseWriter, *http.Request)
}

type api struct {
	logger        *log.Logger
	diceRollerSrv core.DiceRoller
	wallerSrv     core.Wallet
	telemetry     telemetry.Telemetry

	rollCount                 metric.Int64Counter
	doubleValueDistribution   metric.Int64Histogram
	diceSrvResponseTimeMetric metric.Int64Histogram
	diceSrcCalls              metric.Int64Counter
	diceSrvValueDistribution  metric.Int64Histogram
	apiActiveClients          metric.Int64UpDownCounter
	apiMissingPayment         metric.Int64Counter
}

func NewAPI(tel telemetry.Telemetry, logger *log.Logger) API {

	return &api{
		logger:        logger,
		telemetry:     tel,
		diceRollerSrv: core.NewDiceRoller(),
		wallerSrv:     core.NewWallet(tel.Meter()),

		// api metrics
		apiActiveClients:  telemetry.ApiActiveClients(tel.Meter()),
		apiMissingPayment: telemetry.ApiMissingPaymentErrorCount(tel.Meter()),

		// dice metrics
		rollCount:               telemetry.DoubleDiceRollCount(tel.Meter()),
		doubleValueDistribution: telemetry.DoubleDiceValueDistribution(tel.Meter()),

		// dice service metrics
		diceSrvResponseTimeMetric: telemetry.DiceServiceResponseTime(tel.Meter()),
		diceSrcCalls:              telemetry.DiceServiceTotalCalls(tel.Meter()),
		diceSrvValueDistribution:  telemetry.DiceServiceValueDistribution(tel.Meter()),
	}
}

func (a *api) DiceRoller(w http.ResponseWriter, r *http.Request) {

	spanCtx, span := a.telemetry.Tracer().Start(r.Context(), "roll")
	defer span.End()

	// API active clients
	a.addAddActiveClient(spanCtx)
	defer a.removeActiveClient(spanCtx)

	nDices := 2
	valuesChan := make(chan int, nDices)
	for i := 0; i < nDices; i++ {
		go func(_ int) {
			value := a.timedSrvCall(spanCtx)
			valuesChan <- value
		}(i)
	}

	totalValue := 0
	isPair := false
	for i := 0; i < nDices; i++ {
		select {
		case roll := <-valuesChan:
			if totalValue == roll {
				isPair = true
			}
			totalValue += roll
		}
	}

	rollValueAttr := attribute.Int("roll.value", totalValue)
	span.SetAttributes(rollValueAttr)

	a.metricCounter(spanCtx, rollValueAttr)
	a.metricHistogram(spanCtx, totalValue)

	if isPair {
		a.wallerSrv.Deposit()
	}

	if a.wallerSrv.Balance() < 1 {
		a.apiMissingPayment.Add(spanCtx, 1)
		span.SetStatus(http.StatusPaymentRequired, http.StatusText(http.StatusPaymentRequired)+" "+strconv.Itoa(totalValue))
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(http.StatusText(http.StatusPaymentRequired)))
		return
	}

	defer a.wallerSrv.Withdraw()

	span.SetStatus(http.StatusOK, http.StatusText(http.StatusOK))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strconv.Itoa(totalValue)))
}

func (a *api) timedSrvCall(ctx context.Context) int {
	start := time.Now()
	roll := a.diceRollerSrv.Roll()
	elapsed := time.Since(start)

	a.diceSrvResponseTimeMetric.Record(ctx, elapsed.Milliseconds())
	a.diceSrcCalls.Add(ctx, 1)
	a.diceSrvValueDistribution.Record(ctx, int64(roll))

	return roll
}

func (a *api) addAddActiveClient(ctx context.Context) {
	a.logger.Printf("ADD::apiActiveClients: %d", 1)
	a.apiActiveClients.Add(ctx, int64(1))
}

func (a *api) removeActiveClient(ctx context.Context) {
	a.logger.Printf("SUBS::apiActiveClients: %d", -1)
	a.apiActiveClients.Add(ctx, int64(-1))
}

func (a *api) metricCounter(ctx context.Context, rollValueAttr attribute.KeyValue) {
	a.rollCount.Add(ctx, 1, metric.WithAttributes(rollValueAttr))
}

func (a *api) metricHistogram(ctx context.Context, roll int) {
	a.doubleValueDistribution.Record(ctx, int64(roll))
}
