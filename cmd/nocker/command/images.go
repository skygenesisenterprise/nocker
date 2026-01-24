package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewImagesCommand creates the 'images' command
func NewImagesCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "images [OPTIONS] [REPOSITORY[:TAG]]",
		Short:   "List images",
		Long:    `List images`,
		Aliases: []string{"image", "list"},
		Example: "nocker images",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Images command executed")
			// TODO: Implement actual images logic
			return nil
		},
	}
}