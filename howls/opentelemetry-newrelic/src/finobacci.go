package src

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	name = "fibonacci-calculator"
)

var (
	tracer          = otel.Tracer(name)
	meter           = otel.Meter(name)
	logger          = otelslog.NewLogger(name)
	totalCalls      metric.Int64Counter
	activeRequests  metric.Int64Gauge
	calculationTime metric.Float64Histogram
)

type fibonacciResponse struct {
	N       int    `json:"n,omitempty"`
	Result  int    `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

func init() {
	var err error
	totalCalls, err = meter.Int64Counter(
		"fibonacci.TotalCalls",
		metric.WithDescription("Measures the total number of calls to the fibonacci method."),
	)
	if err != nil {
		panic(err)
	}

	activeRequests, err = meter.Int64Gauge("fibonacci.ActiveRequests",
		metric.WithDescription("Measures the number of active requests to the fibonacci method."))
	if err != nil {
		panic(err)
	}

	calculationTime, err = meter.Float64Histogram("fibonacci.CalculationTime",
		metric.WithDescription("Measures the time taken to calculate the fibonacci number.")),
		metric.WithExplicitBucketBoundaries(10, 20, 30, 40, 50)
	if err != nil {
		panic(err)
	}

}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	pv := r.PathValue("n")
	n, err := strconv.Atoi(pv)
	if err != nil {
		logger.Error(err.Error())
		createHttpResponse(w, http.StatusBadRequest, fibonacciResponse{Message: err.Error()})
		return
	}

	// Metric total active requests
	activeRequests.Record(r.Context(), 1)
	defer activeRequests.Record(r.Context(), -1)

	// Metric time to calculate fibonacci
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime).Seconds()
		calculationTime.Record(r.Context(), duration, metric.WithAttributes(attribute.Int("fibonacci.N", n)))
	}()

	result, err := calculateFibonacci(r.Context(), n)
	totalCalls.Add(r.Context(), 1, metric.WithAttributes(attribute.Bool("fibonacci.valid", err == nil)))

	if err != nil {
		createHttpResponse(w, http.StatusBadRequest, fibonacciResponse{Message: err.Error()})
		return
	}

	createHttpResponse(w, http.StatusOK, fibonacciResponse{N: n, Result: result})
}

func calculateFibonacci(ctx context.Context, n int) (int, error) {
	ctx, span := tracer.Start(ctx, "fibonacci")
	defer span.End()

	span.SetAttributes(attribute.Int("fibonacci.n", n))

	if n < 1 || n > 90 {
		err := errors.New("n must be between 1 and 90")
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err, trace.WithStackTrace(true))
		msg := fmt.Sprintf("Failed to compute fib(%d).", n)
		logger.InfoContext(ctx, msg, "fibonacci.n", n)
		return 0, err
	}

	var result = 1
	if n > 2 {
		var a = 0
		var b = 1

		for i := 1; i < n; i++ {
			result = a + b
			a = b
			b = result
		}
	}

	span.SetAttributes(attribute.Int("fibonacci.result", result))
	msg := fmt.Sprintf("Computed fib(%d) = %d.", n, result)
	logger.InfoContext(ctx, msg, "fibonacci.n", n, "fibonacci.result", result)
	return result, nil
}

func createHttpResponse(w http.ResponseWriter, statusCode int, res fibonacciResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
