package logger

import (
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	InitLogger("info", LoggerOpts{Type:LoggerConsole})
}

func TestGetLogger(t *testing.T) {
	InitLogger("info", LoggerOpts{Type:LoggerConsole})

	myLogger := GetLogger("mylogger")
	//myLogger.async = false
	myLogger.Info("Hello, Mylogger")
	time.Sleep(2*time.Second)
}