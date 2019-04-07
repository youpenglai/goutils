package logger

import (
	"time"
	"strings"
	"fmt"
)

const TimeFormatter = "2006-01-02 15:04:05.000"

type LoggerMsg struct {
	Prefix string
	Level int
	Time time.Time
	Caller struct {
		FileName string
		Line int
	}
	Msg string
}

type LoggerFormatter interface {
	Format(colorful bool) string
}

func levelColor(level int) string {
	switch level {
	case LevelDebug:
		return ConsoleBlue
	case LevelInfo:
		return ConsoleGreen
	case LevelWarn:
		return ConsoleYellow
	case LevelError:
		return ConsoleRed
	case LevelAbort:
		return ConsoleMagenta
	}

	return ConsoleReset
}

func (lmsg *LoggerMsg) Format(colorful bool) string {
	prefix := strings.ToUpper(lmsg.Prefix)
	level := levelPrefix[lmsg.Level]
	timeString := lmsg.Time.Format(TimeFormatter)
	caller := fmt.Sprintf("%s:%d:", lmsg.Caller.FileName, lmsg.Caller.Line)
	if colorful {
		prefix = ColorText(prefix, ConsoleCyan)
		level = ColorText(level, levelColor(lmsg.Level))
	}

	return fmt.Sprintf("%s %s %s %s %s\n", prefix, timeString, caller, level, lmsg.Msg)
}
