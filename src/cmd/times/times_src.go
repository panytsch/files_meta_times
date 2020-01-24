package times

import (
	"os"
	"path/filepath"
)

//Writter function to write file
type Writter func(file *File)

//GenerateReport and write it using witter
func GenerateReport(rootPath string, writer Writter) {
	_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		writer(NewFile(path))
		return nil
	})
}
