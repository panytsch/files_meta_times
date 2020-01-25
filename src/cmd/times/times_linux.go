package times

import (
	"os"
	"syscall"
	"time"
)

//FillTimes depends on platform
func FillTimes(fi os.FileInfo, f *File) {
	info := fi.Sys().(*syscall.Stat_t)
	f.aTime = time.Unix(info.Atim.Unix())
	f.mTime = time.Unix(info.Mtim.Unix())
}

//IsAccessDateEnabled checks is last access date enabled in you file system
func IsAccessDateEnabled() bool {
	return true
}
