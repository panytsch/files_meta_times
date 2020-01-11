package export

import (
	"encoding/csv"
	"files_meta_times/src/cmd/times"
	"os"
)

const defaultFileName = "scanResult"

// Exporter makes export of files
type Exporter struct {
	data     [][]string
	fileName string
}

//PrepareFilesToExport prepars slice of files to export
func PrepareFilesToExport(files []string) []*times.File {
	var fileToExport []*times.File
	for _, path := range files {
		fileToExport = append(fileToExport, times.NewFile(path))
	}
	return fileToExport
}

//InitializateFiles init files before export
func (e *Exporter) InitializateFiles(files []*times.File) {
	e.data = [][]string{
		{"Path", "Name", "Last file read time", "Last content changing time", "Last file metadata changing time"},
	}
	for _, file := range files {
		e.data = append(e.data, []string{file.GetPath(), file.GetName(), file.GetLastReadTime().String(), file.GetLastModificationTime().String(), file.GetLastMetaChangingTime().String()})
	}
}

//WriteAll write to file
func (e *Exporter) WriteAll() {
	fileToWrite, err := os.Create(e.GetFileName())
	if err != nil {
		return
	}
	writter := csv.NewWriter(fileToWrite)
	for _, record := range e.data {
		writter.Write(record)
	}
}

//SetFileName set file name
func (e *Exporter) SetFileName(name string) {
	e.fileName = name
}

//GetFileName get file name
func (e *Exporter) GetFileName() string {
	if e.fileName == "" {
		e.fileName = defaultFileName
	}
	return e.fileName
}
