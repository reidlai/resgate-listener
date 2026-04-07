package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage the RES protocol server",
	Long:  "Commands to manage the Technical Analysis Assistant RES protocol server",
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the RES protocol server",
	Long:  "Start the RES protocol server",
	RunE:  runServer,
}

func init() {
	// Add start subcommand
	serverCmd.AddCommand(serverStartCmd)

	// Server flags
	serverStartCmd.Flags().String("res.nats.url", "", "NATS Server URL")

	// Bind flags to Viper (using res-server prefix)
	if err := viper.BindPFlag("res.nats.url", serverStartCmd.Flags().Lookup("res.nats.url")); err != nil {
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
	return nil
}
