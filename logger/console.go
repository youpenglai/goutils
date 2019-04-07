package logger

import (
	"os"
)

const (
	ConsoleBlack = "\x1b[30m"
	ConsoleRed = "\x1b[31m"
	ConsoleGreen = "\x1b[32m"
	ConsoleYellow = "\x1b[33m"
	ConsoleBlue = "\x1b[34m"
	ConsoleMagenta = "\x1b[35m"
	ConsoleCyan ="\x1b[36m"
	ConsoleWhite = "\x1b[37m"

	ConsoleReset = "\x1b[0m"
)

func ColorText(text, color string) string {
	return color + text + ConsoleReset
}

type ConsoleWriter struct {
	writer *os.File
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		writer: os.Stdout,
	}
}

func (cw *ConsoleWriter) Colorful() bool {
	return true
}

func (cw *ConsoleWriter) Write(b []byte) (int, error) {
	return cw.writer.Write(b)
}

func (cw *ConsoleWriter) Flush() {
	cw.writer.Sync()
}