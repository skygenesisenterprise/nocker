package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewBuildCommand creates the 'build' command
func NewBuildCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "build [OPTIONS] PATH | URL | -",
		Short:   "Build an image from a Dockerfile",
		Long:    `Build an image from a Dockerfile`,
		Example: "nocker build -t my-image .",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Build command executed")
			// TODO: Implement actual build logic
			return nil
		},
	}
}