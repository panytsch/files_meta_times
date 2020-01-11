package main

import (
	"files_meta_times/cmd"
)

func main() {
	// fileinfo, _ := os.Stat("/home/panytsch/freelance/edhunter/public/index.php")
	// atime := fileinfo.Sys().(*syscall.Stat_t).Atim
	// fmt.Println(time.Unix(atime.Sec, atime.Nsec))
	cmd.Execute()
}
