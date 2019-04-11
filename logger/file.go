package logger

import (
	"bytes"
	"github.com/youpenglai/goutils/pathtool"
	"os"
	"sync"
	"time"
)

type FileWriter struct {
	mu       sync.Mutex
	filename string

	logStartTime time.Time
	//logCurTime   time.Time
	currentFile  int
	rotate       int
	maxSize      int64
	maxLines     int
	maxFiles     int
	file         *os.File
	writeBytes   int64
	writeLines   int
}

func (fw *FileWriter) openLogFile() error {
	exists, err := pathtool.FileExists(fw.filename)
	if err != nil {
		return err
	}

	return nil
}

func (fw *FileWriter) needRotate() bool {
	if fw.rotate == LogRotateNo {
		return false
	}

	if fw.rotate == LogRotateDay {
		now := time.Now()
		if now.After(fw.logStartTime.Add(24 * time.Hour)) {
			return true
		}
	} else if fw.rotate == LogRotateHour {
		now := time.Now()
		if now.After(fw.logStartTime.Add(time.Hour)) {
			return true
		}
	} else if fw.rotate == LogRotateLines {
		if fw.writeLines > fw.maxLines {
			return true
		}
	}

	return false
}

func (fw *FileWriter) doRotate() error {

	return nil
}

func (fw *FileWriter) Colorful() bool {
	return false
}

func (fw *FileWriter) Write(b []byte) (int, error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if fw.needRotate() {
		fw.doRotate()
	}

	wn, err := fw.file.Write(b)
	if err != nil {
		return 0, err
	}
	fw.writeBytes += int64(wn)
	fw.writeLines += bytes.Count(b, []byte{'\n'})

	return wn, err
}

func (fw *FileWriter) Flush() {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.file.Sync()
}

func NewFileWriter(opts LoggerOpts) LoggerWriter {
	fw := &FileWriter{
		rotate:   opts.Rotate,
		filename: opts.FileName,
		maxSize:  opts.MaxSize,
		maxLines: opts.MaxLines,
	}
	return nil
}

func init() {
	RegisterLoggerWriter(LoggerFile, NewFileWriter)
}
