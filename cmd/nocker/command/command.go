package command

import (
	"github.com/spf13/cobra"
)

// InitCommands initializes all Nocker commands
func InitCommands(rootCmd *cobra.Command) {
	// Common Commands
	rootCmd.AddCommand(NewRunCommand())
	rootCmd.AddCommand(NewExecCommand())
	rootCmd.AddCommand(NewPsCommand())
	rootCmd.AddCommand(NewBuildCommand())
	rootCmd.AddCommand(NewPullCommand())
	rootCmd.AddCommand(NewPushCommand())
	rootCmd.AddCommand(NewImagesCommand())
	rootCmd.AddCommand(NewLoginCommand())
	rootCmd.AddCommand(NewLogoutCommand())
	rootCmd.AddCommand(NewSearchCommand())
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(NewInfoCommand())

	// Management Commands
	rootCmd.AddCommand(NewContainerCommand())
	rootCmd.AddCommand(NewImageCommand())
	rootCmd.AddCommand(NewNetworkCommand())
	rootCmd.AddCommand(NewVolumeCommand())
	rootCmd.AddCommand(NewSystemCommand())

	// Additional Commands
	rootCmd.AddCommand(NewAttachCommand())
	rootCmd.AddCommand(NewCommitCommand())
	rootCmd.AddCommand(NewCpCommand())
	rootCmd.AddCommand(NewCreateCommand())
	rootCmd.AddCommand(NewDiffCommand())
	rootCmd.AddCommand(NewEventsCommand())
	rootCmd.AddCommand(NewExportCommand())
	rootCmd.AddCommand(NewHistoryCommand())
	rootCmd.AddCommand(NewImportCommand())
	rootCmd.AddCommand(NewInspectCommand())
	rootCmd.AddCommand(NewKillCommand())
	rootCmd.AddCommand(NewLoadCommand())
	rootCmd.AddCommand(NewLogsCommand())
	rootCmd.AddCommand(NewPauseCommand())
	rootCmd.AddCommand(NewPortCommand())
	rootCmd.AddCommand(NewRenameCommand())
	rootCmd.AddCommand(NewRestartCommand())
	rootCmd.AddCommand(NewRmCommand())
	rootCmd.AddCommand(NewRmiCommand())
	rootCmd.AddCommand(NewSaveCommand())
	rootCmd.AddCommand(NewStartCommand())
	rootCmd.AddCommand(NewStatsCommand())
	rootCmd.AddCommand(NewStopCommand())
	rootCmd.AddCommand(NewTagCommand())
	rootCmd.AddCommand(NewTopCommand())
	rootCmd.AddCommand(NewUnpauseCommand())
	rootCmd.AddCommand(NewUpdateCommand())
	rootCmd.AddCommand(NewWaitCommand())
}