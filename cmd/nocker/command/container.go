package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewContainerCommand creates the 'container' management command
func NewContainerCommand() *cobra.Command {
	var containerCmd = &cobra.Command{
		Use:     "container",
		Short:   "Manage containers",
		Long:    `Manage containers`,
		Aliases: []string{"containers"},
	}

	// Subcommands
	containerCmd.AddCommand(&cobra.Command{
		Use:     "ls [OPTIONS]",
		Short:   "List containers",
		Long:    `List containers`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Container ls command executed")
			return nil
		},
	})

	containerCmd.AddCommand(&cobra.Command{
		Use:     "prune [OPTIONS]",
		Short:   "Remove all stopped containers",
		Long:    `Remove all stopped containers`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Container prune command executed")
			return nil
		},
	})

	return containerCmd
}