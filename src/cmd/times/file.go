package times

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// File - object with file data
type File struct {
	path    string
	ext     string
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

// GetExt returns file extension
func (f *File) GetExt() string {
	return filepath.Ext(f.path)
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
	fileInfo, _ := os.Stat(f.path)
	FillTimes(fileInfo, f)
	f.initted = true
}

// GetPath returns absolute file path
func (f *File) GetPath() string {
	return f.path
}

func (f *File) GetRecord() []string {
	return []string{
		f.GetPath(),
		f.GetName(),
		f.GetExt(),
		f.GetLastReadTime().String(),
		f.GetLastModificationTime().String(),
	}
}

func (f *File) GetFormattedRecord() []string {
	return []string{
		f.GetPath(),
		f.GetName(),
		f.GetExt(),
		format(f.GetLastReadTime()),
		format(f.GetLastModificationTime()),
	}
}

func format(t time.Time) string {
	result := strconv.Itoa(t.Year()) + "-"
	result += strconv.Itoa(int(t.Month())) + "-"
	result += strconv.Itoa(t.Day()) + " "
	result += strconv.Itoa(t.Hour()) + ":"
	result += strconv.Itoa(t.Minute()) + ":"
	result += strconv.Itoa(t.Second())
	return result
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
