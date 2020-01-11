package times

import (
	"os"
	"path/filepath"
)

//GetFilesList return all files in folder
func GetFilesList(rootPath string) (result []string) {
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			absoluteFilePath, er := filepath.Abs(path)
			if er != nil {
				return er
			}
			result = append(result, absoluteFilePath)
		}
		return nil
	})
	return result
}
