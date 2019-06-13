package logger

import (
	"testing"
	"time"
)

func TestNewFileWriter(t *testing.T) {
	fw := NewFileWriter(LoggerOpts{
		FileName: "tmp/test.log",
		Type: LoggerFile,
		Rotate:LogRotateHour,
	})
	_, err := fw.Write([]byte("Hello, Test log" + time.Now().Format(time.ANSIC)))
	if err != nil {
		t.Error(err)
		return
	}
	fw.Flush()
}
