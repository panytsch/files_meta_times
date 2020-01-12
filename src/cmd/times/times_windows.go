package times

import (
	"os"
	"syscall"
	"time"
)

//FillTimes depends on platform
func FillTimes(fi os.FileInfo, f *File) {
	info := fi.Sys().(*syscall.Win32FileAttributeData)
	f.aTime = time.Unix(0, info.LastAccessTime.Nanoseconds())
	f.mTime = time.Unix(0, info.LastWriteTime.Nanoseconds())
}
