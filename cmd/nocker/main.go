package main

import (
	"fmt"
	"os"

	"github.com/skygenesisenterprise/nocker/cmd/command"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:              "nocker",
		Short:            "A self-sufficient runtime for containers",
		Long:             `Nocker - A modern container runtime inspired by Docker`,
		Version:          fmt.Sprintf("%s (%s, %s)", version, commit, date),
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
	}

	// Add global flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "Location of client config files")
	rootCmd.PersistentFlags().StringP("context", "x", "", "Name of the context to use")
	rootCmd.PersistentFlags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringP("host", "H", "", "Daemon socket to connect to")
	rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Set the logging level")

	// Initialize commands
	command.InitCommands(rootCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}