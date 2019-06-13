package logger

import (
	"testing"
	"time"
)

func TestNewTextFormatter(t *testing.T) {
	t.Log("success")
}

func TestTextFormatter_Format(t *testing.T) {
	msg := LoggerMsg{
		Prefix: "Test",
		Level: LevelDebug,
		Time: time.Date(2000, 1,1,0,0,0,0, time.Local),
		Caller: struct {
			FileName string
			Line     int
		}{FileName: "test.go", Line: 10},
		Msg: "Debug message here",
	}

	t.Log(msg.Format(true))

	t.Log("success")
}