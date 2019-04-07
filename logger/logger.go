package logger

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	LevelAbort = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug

	LoggerConsole = "console"
	LoggerFile    = "file"

	LogRotateNo = 0
	LogRotateLines = 1
	LogRotateDay = 2
	LogRotateHour = 3
)

var (
	levelPrefix []string = []string{"[ABORT]", "[ERROR]", "[WARN]", "[INFO]", "[DEBUG]"}
	levelNames  []string = []string{"error", "info"}
)

type LoggerWriter interface {
	Colorful() bool
	Write([]byte) (int, error)
	Flush()
}

type LoggerOpts struct {
	Type string
	Prefix string
	Level  int
	Rotate int
	MaxLines int
}

var loggerWriters map[string]LoggerWriter

type Logger struct {
	name   string
	async  bool
	level  int
	prefix string

	buff    chan *LoggerMsg
	writers []LoggerWriter
	mu      sync.Mutex
}

func NewLogger(opts LoggerOpts) *Logger {
	logger := &Logger{}

	// start write job
	go func() {
		for {
			select {
			case msg := <-logger.buff:
				logger.mu.Lock()
				defer logger.mu.Unlock()
				logger.writeTask(msg)
			}
		}
	}()

	return logger
}

func (l *Logger) writeTask(msg *LoggerMsg) {
	var m string
	var colorful bool
	for _, writer := range l.writers {
		if m == "" || colorful != writer.Colorful() {
			colorful = writer.Colorful()
			m = msg.Format(writer.Colorful())
		}
		writer.Write([]byte(m))
	}
}

var msgPool *sync.Pool

func (l *Logger) allocMsg() *LoggerMsg {
	v := msgPool.Get()
	msg, _ := v.(*LoggerMsg)
	msg.Msg = ""
	msg.Caller.Line = -1
	msg.Caller.FileName = ""
	msg.Prefix = ""

	return msg
}

func (l *Logger) freeMsg(msg *LoggerMsg) {
	msgPool.Put(msg)
}

func (l *Logger) Debug(f string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}
	l.writeMsg(f, v)
}

func (l *Logger) Info(f string, v ...interface{}) {
	if LevelInfo > l.level {
		return
	}
	l.writeMsg(f, v)
}

func (l *Logger) Warn(f string, v ...interface{}) {
	if LevelWarn > l.level {
		return
	}
	l.writeMsg(f, v)
}

func (l *Logger) Error(f string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	l.writeMsg(f, v)
}

func (l *Logger) Abort(f string, v ...interface{}) {
	if LevelAbort > l.level {
		return
	}
	l.writeMsg(f, v)
}

func (l *Logger) writeMsg(f string, v ...interface{}) {
	rawMsg := fmt.Sprintf(f, v)
	msg := l.allocMsg()

	msg.Msg = rawMsg
	msg.Prefix = l.prefix
	msg.Time = time.Now()
	msg.Level = l.level

	_, file, line, _ := runtime.Caller(2)
	msg.Caller.FileName = file
	msg.Caller.Line = line

	l.buff <- msg
	if !l.async {
		// flush
	}
}

func init() {
	// initialize pool
	msgPool = &sync.Pool{
		New: func() interface{} {
			return &LoggerMsg{}
		},
	}
}

func InitLogger() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	return &Logger{}
}
