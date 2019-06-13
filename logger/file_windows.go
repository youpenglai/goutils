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

	info, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return t, errors.New("can not get file creation time")
	}
	t = time.Unix(0, info.CreationTime.Nanoseconds())

	return
}
