package times

import (
	"os"
	"path/filepath"
)

func scanDir(rootPath string) (result []string) {
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		result = append(result, path)
		return nil
	})
	return result
}
