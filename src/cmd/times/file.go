package times

import (
	"os"
	"syscall"
	"time"
)

// File - object with file data
type File struct {
	path    string
	aTime   time.Time // last file read time
	mTime   time.Time // time when content changed
	cTime   time.Time // last file meta data changed
	initted bool
}

// NewFile returns new instance of file
func NewFile(path string) *File {
	file := &File{}
	file.path = path
	file.init()
	return file
}

//GetLastReadTime in unix time
func (f *File) GetLastReadTime() time.Time {
	return f.aTime
}

//GetLastModificationTime in unix time
func (f *File) GetLastModificationTime() time.Time {
	return f.mTime
}

//GetLastMetaChangingTime in unix time
func (f *File) GetLastMetaChangingTime() time.Time {
	return f.cTime
}

func (f *File) init() {
	if f.initted {
		return
	}
	fileinfo, _ := os.Stat(f.path)
	info := fileinfo.Sys().(*syscall.Stat_t)
	f.aTime = time.Unix(info.Atim.Unix())
	f.mTime = time.Unix(info.Mtim.Unix())
	f.cTime = time.Unix(info.Ctim.Unix())
	f.initted = true
}
