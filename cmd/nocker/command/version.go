package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCommand creates the 'version' command
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Show the Nocker version information",
		Long:    `Show the Nocker version information`,
		Example: "nocker version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Nocker version 0.1.0")
			fmt.Println("Build: dev")
			fmt.Println("Commit: none")
			fmt.Println("Date: unknown")
			return nil
		},
	}
}