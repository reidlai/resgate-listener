package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/nats-io/nats.go"
	example "github.com/reidlai/resgate-listener/example/example_handlers"
	"github.com/reidlai/resgate-listener/internal/otelutils"
	"github.com/reidlai/resgate-listener/pkg/listener"
	rmh "github.com/reidlai/resgate-listener/pkg/resgate_message_handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	otelslog "go.opentelemetry.io/contrib/bridges/otelslog"
	otel "go.opentelemetry.io/otel"
	otel_stdouttrace "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	otel_global "go.opentelemetry.io/otel/log/global"
	oltel_propagation "go.opentelemetry.io/otel/propagation"
	otel_log "go.opentelemetry.io/otel/sdk/log"
	otel_resource "go.opentelemetry.io/otel/sdk/resource"
	otel_sdktrace "go.opentelemetry.io/otel/sdk/trace"
	otel_semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Manage the RES protocol server",
	Long:  "Commands to manage the Technical Analysis Assistant RES protocol server",
	RunE: startServer,
}

func init() {
	// Add start subcommand

	// Server flags
	startCmd.Flags().String("resgate.nats-url", "", "NATS Server URL")


	// Bind flags to Viper
	if err := viper.BindPFlag("resgate.nats-url", startCmd.Flags().Lookup("resgate.nats-url")); err != nil {
		panic(err)
	}

	// Set Default values in Viper
	viper.SetDefault("resgate.nats-url", "nats://localhost:4222")

	// Explicit BindEnv
	_ = viper.BindEnv("resgate.nats-url")
}

func initOTel() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := otel_resource.New(ctx,
		otel_resource.WithAttributes(
			otel_semconv.ServiceName("resgate-listener-logger"),
		),
	)
	if err != nil {
		return nil, err
	}

	// 1. Tracing Setup
	traceExporter, err := otel_stdouttrace.New()
	if err != nil {
		return nil, err
	}

	tp := otel_sdktrace.NewTracerProvider(
		otel_sdktrace.WithSampler(otel_sdktrace.AlwaysSample()),
		otel_sdktrace.WithResource(res),
		otel_sdktrace.WithBatcher(traceExporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(oltel_propagation.NewCompositeTextMapPropagator(oltel_propagation.TraceContext{}, oltel_propagation.Baggage{}))

	// 2. Logging Setup
	var logExporter otel_log.Exporter
	format := viper.GetString("server.log-format")

	if format == "json" {
		logExporter = otelutils.NewCompactExporter(os.Stdout)
	} else {
		// Use the same compact exporter for now to satisfy "Standard" but we could 
		// use a text exporter if needed.
		logExporter = otelutils.NewCompactExporter(os.Stdout)
	}

	lp := otel_log.NewLoggerProvider(
		otel_log.WithResource(res),
		otel_log.WithProcessor(otel_log.NewSimpleProcessor(logExporter)),
	)
	otel_global.SetLoggerProvider(lp)

	return func(c context.Context) error {
		_ = tp.Shutdown(c)
		_ = lp.Shutdown(c)
		return nil
	}, nil
}

func startServer(cmd *cobra.Command, args []string) error {
	// Initialize OTel Trace and Log SDKs
	shutdown, err := initOTel()
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to initialize OTel", "error", err)
	} else {
		defer shutdown(context.Background())
	}

	// Initialize slog with official OTel bridge
	// This delegates all slog calls to the OTel LoggerProvider initialized above.
	logger := slog.New(otelslog.NewHandler("resgate-listener-logger"))
	slog.SetDefault(logger)

	natsURL := viper.GetString("resgate.nats-url")
	logFormat := viper.GetString("server.log-format")
	slog.InfoContext(context.Background(), "Starting dummy server...", "resgate.nats-url", natsURL, "server.log-format", logFormat)

	// Create NATS connection
	nc, err := nats.Connect(natsURL)
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to connect to NATS", "error", err)
		return err
	}
	defer nc.Close()

	exampleHandler := &example.ExampleHandler{
		BaseResgateMessageHandler: &rmh.BaseResgateMessageHandler{},
	}

	// Create handler mapping
	handlers := map[string]rmh.ResgateMessageHandler{
		"get.examples":    exampleHandler,
		"get.example.*": exampleHandler,
	}

	listenerInstance, err := listener.NewResgateListener(nc, handlers)
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to create listener", "error", err)
		return err
	}
	defer listenerInstance.Close()

	if err := listenerInstance.Listen(); err != nil {
		slog.ErrorContext(context.Background(), "Failed to start listening", "error", err)
		return err
	}
	
	return nil
}
