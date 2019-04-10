package logger

import (
	"fmt"
	"runtime"
	"strings"
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

	LogRotateNo    = 0
	LogRotateLines = 1
	LogRotateDay   = 2
	LogRotateHour  = 3
)

var (
	levelPrefix []string = []string{"[ABORT]", "[ERROR]", "[WARN]", "[INFO]", "[DEBUG]"}
	levelNames  []string = []string{"error", "info"}
)

func getLevelByName(name string) int {
	switch strings.ToLower(name) {
	case "abort":
		return LevelAbort
	case "error":
		return LevelError
	case "warn", "warning":
		return LevelWarn
	case "info", "information":
		return LevelInfo
	case "debug", "dbg":
		return LevelDebug
	default:
		return LevelDebug
	}
}

type LoggerWriter interface {
	Colorful() bool
	Write([]byte) (int, error)
	Flush()
}

type LoggerOpts struct {
	Type     string `json:"type"`
	FileName string `json:"fileName"`
	Rotate   int    `json:"rotate,omitempty"`
	MaxLines int    `json:"maxLines,omitempty"`
	MaxSize  int64  `json:"maxSize,omitempty"`
}

var loggerWriters map[string]LoggerWriterCreateFunc

type LoggerWriterCreateFunc func(opts LoggerOpts) LoggerWriter

func RegisterLoggerWriter(name string, loggerCreateFunc LoggerWriterCreateFunc) {
	loggerWriters[name] = loggerCreateFunc
}

func getLoggerWriterCreateFunc(name string) LoggerWriterCreateFunc{
	w, e := loggerWriters[name]
	if e {
		return w
	}
	name = LoggerConsole
	w, _ = loggerWriters[name]
	return w
}

type Logger struct {
	async  bool
	level  int
	prefix string

	buff    chan *LoggerMsg
}

const defaultBuffSize = 100

func NewLogger(prefix, level string, async bool) *Logger {
	logger := &Logger{
		async:  async,
		level:  getLevelByName(level),
		prefix: prefix,
	}

	return logger
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

func (l *Logger) SetLevel(level int) {
	l.level = level
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

	writeLogMsg(msg, l.async)
}


var loggers = struct {
	mu            sync.Mutex
	level         string
	loggers       map[string]*Logger
	defaultLogger *Logger
	writers       []LoggerWriter
	buffers       chan *LoggerMsg
	flush         chan int
}{}

func writeLogMsg(msg *LoggerMsg, async bool) {
	if loggers.buffers == nil {
		return
	}

	loggers.buffers <- msg
	if async {
		loggers.flush <- 1
	}
}

func writeTasks() {
	flush := false
	flushTicker := time.NewTicker(time.Second)
	for {
		select {
		case <- flushTicker.C:
			for _, writer := range loggers.writers {
				writer.Flush()
			}
		case <- loggers.flush:
			flush = true
			continue
		case msg := <- loggers.buffers:
			for _, writer := range loggers.writers {
				writer.Write([]byte(msg.Format(writer.Colorful())))
				if flush && len(loggers.buffers) == 0 {
					writer.Flush()
					flush = false
				}
			}
		}
	}
}

func init() {
	// initialize pool
	msgPool = &sync.Pool{
		New: func() interface{} {
			return &LoggerMsg{}
		},
	}

	go writeTasks()
}

func InitLogger(level string, opts ...LoggerOpts) error {
	loggers.mu.Lock()
	defer loggers.mu.Unlock()

	loggers.buffers = make(chan *LoggerMsg, defaultBuffSize)
	loggers.flush = make(chan int)

	loggers.level = level
	if len(opts) == 0 {
		opts = append(opts, LoggerOpts{Type:LoggerConsole})
	}
	for _, opt := range opts {
		createFunc := getLoggerWriterCreateFunc(opt.Type)
		writer := createFunc(opt)
		loggers.writers = append(loggers.writers, writer)
	}

	return nil
}

func GetLogger(prefix string) *Logger {
	loggers.mu.Lock()
	defer loggers.mu.Unlock()

	if prefix == "" {
		return loggers.defaultLogger
	}

	logger, exist := loggers.loggers[prefix]
	if !exist {
		logger = NewLogger(prefix, loggers.level, true)
		loggers.loggers[prefix] = logger
	}

	return logger
}
