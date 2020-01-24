package export

import (
	"encoding/csv"
	"files_meta_times/src/cmd/times"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const defaultFileName = "scanResult"

// Exporter makes export of files
type Exporter struct {
	fileName            string
	exportFile          *os.File
	isExportFileInitted bool
	csvWriter           *csv.Writer
}

func (e *Exporter) Write(file *times.File) {
	var err error
	if e.isExportFileInitted == false {
		err = e.initExportFile()
		if err != nil {
			return
		}
		e.csvWriter = csv.NewWriter(e.exportFile)
		_ = e.csvWriter.Write([]string{"Path", "Name", "Extension", "Last file read time", "Last content changing time"})
	}
	er := e.csvWriter.Write(file.GetRecord())
	if er != nil {
		log.Fatal("error during writing", er.Error())
	}
}

func (e *Exporter) Flush() {
	e.csvWriter.Flush()
	e.exportFile.Close()
}

func (e *Exporter) initExportFile() error {
	var err error
	e.exportFile, err = os.Create(e.GetFileName())
	if err == nil {
		e.isExportFileInitted = true
	}
	return err
}

//SetFileName set file name
func (e *Exporter) SetFileName(name string) {
	e.fileName = name + ".csv"
}

//GetFileName get file name
func (e *Exporter) GetFileName() string {
	if e.fileName == "" {
		e.SetFileName(defaultFileName)
	}
	return e.fileName
}

//CommandExportRun basic export command
func CommandExportRun(_ *cobra.Command, args []string) {
	exporter := &Exporter{}
	if len(args) >= 2 {
		exporter.SetFileName(args[1])
	}
	times.GenerateReport(args[0], exporter.Write)
	defer exporter.Flush()
}
