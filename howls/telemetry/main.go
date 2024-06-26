package main

import (
	"context"
	"errors"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/config/otelcfg"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/presentation"
	"github.com/TheBigBadWolfClub/go-lab/howls/telemetry/internal/shared/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := otelcfg.SetupOTelSDKHTTP(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8091",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	newTelemetry := telemetry.NewTelemetry("dice-roller")

	// Register handlers.
	api := presentation.NewAPI(newTelemetry, logger)
	handleFunc("/api/v1/rolldice", api.DiceRoller)

	generatorAPI := presentation.NewGeneratorAPI(logger)
	handleFunc("/api/v1/dealer", generatorAPI.GenerateDiceRollers)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
