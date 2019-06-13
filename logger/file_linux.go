package logger

import (
	"time"
	"os"
	"syscall"
	"errors"
)

func getFileCreatedTime(f string) (t time.Time, err error) {
	var fi os.FileInfo
	if fi, err = os.Stat(f); err != nil {
		return
	}

	info, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return t, errors.New("can not get file creation time")
	}
	t = time.Unix(info.Ctim.Sec, info.Ctim.Nsec)

	return
}
