package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewPsCommand creates the 'ps' command
func NewPsCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "ps [OPTIONS]",
		Short:   "List containers",
		Long:    `List containers`,
		Aliases: []string{"container", "containers"},
		Example: "nocker ps -a",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Ps command executed")
			// TODO: Implement actual ps logic
			return nil
		},
	}
}