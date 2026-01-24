package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewExecCommand creates the 'exec' command
func NewExecCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "exec [OPTIONS] CONTAINER COMMAND [ARG...]",
		Short:   "Execute a command in a running container",
		Long:    `Execute a command in a running container`,
		Example: "nocker exec -it mycontainer bash",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Exec command executed")
			// TODO: Implement actual exec logic
			return nil
		},
	}
}