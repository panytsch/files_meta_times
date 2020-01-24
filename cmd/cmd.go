package cmd

import (
	"files_meta_times/src/cmd/export"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCmd)
}

// Execute - executing command
func Execute() {
	_ = rootCmd.Execute()
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Get file times from metadata",
	Args:  cobra.MinimumNArgs(1),
	Run:   export.CommandExportRun,
}

var rootCmd = &cobra.Command{
	Use: "app",
}
