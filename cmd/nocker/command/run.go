package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewRunCommand creates the 'run' command
func NewRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
		Short:   "Create and run a new container from an image",
		Long:    `Create and run a new container from an image`,
		Example: "nocker run -it ubuntu bash",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Run command executed")
			// TODO: Implement actual run logic
			return nil
		},
	}
}