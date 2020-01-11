package cmd

import (
	"files_meta_times/src/cmd/export"
	"files_meta_times/src/cmd/times"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(timesCmd)
	rootCmd.AddCommand(filesListCmd)
}

// Execute - executing command
func Execute() {
	rootCmd.Execute()
}

var timesCmd = &cobra.Command{
	Use:   "times",
	Short: "Get file times from metadata",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You should provide path to folder")
			return
		}
		files := export.PrepareFilesToExport(times.GetFilesList(args[0]))
		exporter := &export.Exporter{}
		if len(args) >= 2 {
			exporter.SetFileName(args[1])
		}
		exporter.InitializateFiles(files)
		exporter.WriteAll()
	},
}

var filesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get files list recurcively",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command")
	},
}

var rootCmd = &cobra.Command{
	Use: "app",
}
