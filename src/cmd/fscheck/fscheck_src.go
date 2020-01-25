package fscheck

import (
	"files_meta_times/src/cmd/times"
	"github.com/spf13/cobra"
	"log"
)

func CommandFSCheckRun(_ *cobra.Command, _ []string) {
	if !times.IsAccessDateEnabled() {
		log.Println("FALSE! Last access time isn't enabled in your File system. Please enable it")
		return
	}
	log.Println("OK! Last access time is enabled")
}
