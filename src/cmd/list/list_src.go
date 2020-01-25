package list

import (
	"files_meta_times/src/cmd/times"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const Separator = "\t"

func PrintFile(f *times.File) {
	line := strings.Join(GetFormattedRecord(f), Separator)
	PrintString(line)
}

func PrintString(line string) {
	_, err := os.Stderr.WriteString(line + "\n")
	if err != nil {
		log.Fatalln("Cannot show you anything. Sorry")
	}
}

func CommandListRun(_ *cobra.Command, args []string) {
	if !times.IsAccessDateEnabled() {
		log.Println("WARNING: Last access date is disabled in your file system. please enable it. Report will be useless")
	}
	PrintString(strings.Join(GetHeader(), Separator))
	times.GenerateReport(args[0], PrintFile)
}

func GetHeader() []string {
	return []string{
		"Path",
		"Name",
		"Extension",
		"Last file read time",
		"Last content changing time",
	}
}

func GetFormattedRecord(f *times.File) []string {
	return []string{
		f.GetPath(),
		f.GetName(),
		f.GetExt(),
		format(f.GetLastReadTime()),
		format(f.GetLastModificationTime()),
	}
}

func format(t time.Time) string {
	result := zeroConcat(t.Year()) + "-"
	result += zeroConcat(int(t.Month())) + "-"
	result += zeroConcat(t.Day()) + " "
	result += zeroConcat(t.Hour()) + ":"
	result += zeroConcat(t.Minute()) + ":"
	result += zeroConcat(t.Second())
	return result
}

func zeroConcat(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}
