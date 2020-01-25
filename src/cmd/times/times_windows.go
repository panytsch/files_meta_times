package times

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

//FillTimes depends on platform
func FillTimes(fi os.FileInfo, f *File) {
	info := fi.Sys().(*syscall.Win32FileAttributeData)
	f.aTime = time.Unix(0, info.LastAccessTime.Nanoseconds())
	f.mTime = time.Unix(0, info.LastWriteTime.Nanoseconds())
}

//IsAccessDateEnabled checks is last access date enabled in you file system
func IsAccessDateEnabled() bool {
	out, err := exec.Command("fsutil", "behavior", "query", "disablelastaccess").Output()
	if err != nil {
		log.Println("Cannot check your file system support last access date")
		return false
	}
	return strings.Contains(string(out), "DisableLastAccess = 0")
}
