package cmd

import (
	"files_meta_times/src/cmd/export"
	"files_meta_times/src/cmd/fscheck"
	"files_meta_times/src/cmd/list"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(fsCheckCmd)
}

// Execute - executing command
func Execute() {
	_ = rootCmd.Execute()
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Get file times from metadata and write it to file",
	Args:  cobra.MinimumNArgs(1),
	Run:   export.CommandExportRun,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get file times from metadata and show it",
	Args:  cobra.MinimumNArgs(1),
	Run:   list.CommandListRun,
}

var fsCheckCmd = &cobra.Command{
	Use:   "fscheck",
	Short: "Check your file system for supporting last access date",
	Run:   fscheck.CommandFSCheckRun,
}

var rootCmd = &cobra.Command{
	Use: "app",
}
