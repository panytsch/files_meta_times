package times

import (
	"os"
	"path/filepath"
	"time"
)

// File - object with file data
type File struct {
	path    string
	aTime   time.Time // last file read time
	mTime   time.Time // time when content changed
	initted bool
}

// NewFile returns new instance of file
func NewFile(path string) *File {
	file := &File{}
	file.path = path
	file.init()
	return file
}

// GetName returns file name
func (f *File) GetName() string {
	return filepath.Base(f.path)
}

//GetLastReadTime in unix time
func (f *File) GetLastReadTime() time.Time {
	return f.aTime
}

//GetLastModificationTime in unix time
func (f *File) GetLastModificationTime() time.Time {
	return f.mTime
}

func (f *File) init() {
	if f.initted {
		return
	}
	fileinfo, _ := os.Stat(f.path)
	FillTimes(fileinfo, f)
	f.initted = true
}

// GetPath returns absolute file path
func (f *File) GetPath() string {
	return f.path
}
