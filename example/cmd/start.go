package cmd

import (
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/reidlai/resgate-listener/pkg/listener"
	rmh "github.com/reidlai/resgate-listener/pkg/resgate_message_handler"
	example "github.com/reidlai/resgate-listener/example/example_handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Manage the RES protocol server",
	Long:  "Commands to manage the Technical Analysis Assistant RES protocol server",
	RunE: runServer,
}


func init() {
	// Add start subcommand

	// Server flags
	startCmd.Flags().String("res.nats.url", "", "NATS Server URL")

	// Bind flags to Viper (using res-server prefix)
	if err := viper.BindPFlag("res.nats.url", startCmd.Flags().Lookup("res.nats.url")); err != nil {
		panic(err)
	}

	// Set Default values in Viper (instead of Cobra) to allow ENV overrides
	viper.SetDefault("res.nats.url", "nats://localhost:4222")

	// Explicit BindEnv for nats-url
	_ = viper.BindEnv("res.nats.url")
}

func runServer(cmd *cobra.Command, args []string) error {
	natsURL := viper.GetString("res.nats.url")
	slog.Info("Starting dummy server...", "nats.url", natsURL)

	// Create NATS connection
	nc, err := nats.Connect(natsURL)
	if err != nil {
		slog.Error("Failed to connect to NATS", "error", err)
		return err
	}
	defer nc.Close()

	exampleHandler := &example.ExampleHandler{
		BaseResgateMessageHandler: &rmh.BaseResgateMessageHandler{},
	}

	// Create handler mapping
	handlers := map[string]rmh.ResgateMessageHandler{
		"get.examples":    exampleHandler,
		"get.resource.*": exampleHandler,
	}

	listenerInstance, err := listener.NewResgateListener(nc, handlers)
	if err != nil {
		slog.Error("Failed to create listener", "error", err)
		return err
	}
	defer listenerInstance.Close()

	if err := listenerInstance.Listen(); err != nil {
		slog.Error("Failed to start listening", "error", err)
		return err
	}
	
	return nil
}
