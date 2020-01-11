package times

import (
	"os"
	"path/filepath"
	"time"
)

// File - object with file data
type File struct {
	Path  string
	Atime time.Time
}

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
