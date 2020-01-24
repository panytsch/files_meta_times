package list

import (
	"files_meta_times/src/cmd/times"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const Separator = "\t"

func PrintFile(f *times.File) {
	line := strings.Join(f.GetFormattedRecord(), Separator)
	PrintString(line)
}

func PrintString(line string) {
	_, err := os.Stderr.WriteString(line + "\n")
	if err != nil {
		log.Fatalln("Cannot show you anything. Sorry")
	}
}

func CommandListRun(_ *cobra.Command, args []string) {
	PrintString(strings.Join(times.GetHeader(), Separator))
	times.GenerateReport(args[0], PrintFile)
}
