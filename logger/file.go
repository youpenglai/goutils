package logger

import (
	"bytes"
	"github.com/youpenglai/goutils/pathtool"
	"os"
	"sync"
	"time"
	"path/filepath"
	"strings"
	"fmt"
)

type FileWriter struct {
	mu       sync.Mutex
	filename string
	prefixs  []string

	logStartTime time.Time
	//logCurTime   time.Time
	//currentFile  int
	rotate       int
	maxSize      int64
	maxLines     int
	maxFiles     int
	file         *os.File
	writeBytes   int64
	writeLines   int
}

func (fw *FileWriter) openLogFile() error {
	err := os.MkdirAll(filepath.Dir(fw.filename), os.ModeDir)
	// TODO: err process

	fw.file, err = os.OpenFile(fw.filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	t, _ := getFileCreatedTime(fw.filename)
	if fw.rotate == LogRotateDaily {
		fw.logStartTime = t.Truncate(24 * time.Hour)
	} else if fw.rotate == LogRotateHour {
		fw.logStartTime = t.Truncate(time.Hour)
	}

	go func() {
		// count lines
		fi, _ := os.Stat(fw.filename)
		fw.writeBytes = fi.Size()
		buff := make([]byte, 32768)

		lines := 0
		var pos int64
		for pos = 0; ;pos += 32768 {
			n, e := fw.file.ReadAt(buff, pos)
			if e != nil {
				return
			}
			lines += bytes.Count(buff, []byte{'\n'})
			if n < 32768 {
				break
			}
		}
		fw.mu.Lock()
		defer fw.mu.Unlock()
		fw.writeLines += lines
	}()

	return err
}

func (fw *FileWriter) needRotate() bool {
	if fw.rotate == LogRotateNo {
		return false
	}

	if fw.rotate == LogRotateDaily {
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

const (
	dailyFmt = "2006-01-02"
	hourFmt="20060102-15"
	linesFormat = "20060102-1504"
)

func newFileName(oldName, additional string) string {
	suffix := filepath.Ext(oldName)
	name := strings.TrimSuffix(filepath.Base(oldName), suffix)
	return fmt.Sprintf("%s.%s%s", filepath.Join(filepath.Dir(oldName), name), additional, suffix)
}

func (fw *FileWriter) getLogs() ([]string, error) {
	basename := filepath.Base(fw.filename)

	return pathtool.GetDirFiles(filepath.Dir(fw.filename), func(filename string) bool {
		base := filepath.Base(filename)
		names := strings.Split(base, ".")
		return basename == names[0] + "." + names[2]
	})
}

func (fw *FileWriter) removeOldestLog() error {
	if fw.maxFiles == 0 {
		return nil
	}

	logFiles, err := fw.getLogs()
	if err != nil {
		return err
	}

	if len(logFiles) > fw.maxFiles {
		return nil
	}

	var oldestFile string
	oldestTime := time.Now()
	for _, f := range logFiles {
		t, e := getFileCreatedTime(f)
		if e != nil {
			continue
		}
		if t.Before(oldestTime) {
			oldestFile = f
		}
	}

	if len(oldestFile) == 0 {
		return nil
	}

	os.Remove(oldestFile)
	return nil
}

func (fw *FileWriter) rotateDaily() error {
	err := os.Rename(fw.filename, newFileName(fw.filename, time.Now().Format(dailyFmt)))
	if err != nil {
		return err
	}

	return fw.removeOldestLog()
}

func (fw *FileWriter) rotateHour() error {
	err := os.Rename(fw.filename, newFileName(fw.filename, time.Now().Format(hourFmt)))
	if err != nil {
		return err
	}

	return fw.removeOldestLog()
}

func (fw *FileWriter) rotateLines() error {
	// 为了方便处理，文件重命名依然使用时间戳，精确到分
	// 序号还是难以处理的
	err := os.Rename(fw.filename, newFileName(fw.filename, time.Now().Format(linesFormat)))
	if err != nil {
		return err
	}

	return fw.removeOldestLog()
}

func (fw *FileWriter) doRotate() error {
	var err error
	fw.file.Close()
	if fw.rotate == LogRotateDaily {
		err = fw.rotateDaily()
	} else if fw.rotate == LogRotateHour {
		err = fw.rotateHour()
	} else if fw.rotate == LogRotateLines {
		err = fw.rotateLines()
	}
	if err != nil {
		return err
	}

	return fw.openLogFile()
}

func (fw *FileWriter) MatchPrefix(prefix string) bool {
	if len(fw.prefixs) == 0 {
		return true
	}

	for _, p := range fw.prefixs {
		if p == prefix {
			return true
		}
	}

	return false
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
		prefixs: opts.Prefixs,
	}

	// 如果按行数，设置时最小不小于1000，如果小于1000，默认为1000000行
	if fw.rotate == LogRotateLines && fw.maxLines < 1000 {
		fw.maxLines = 1000000
	}

	if err := fw.openLogFile(); err != nil {
		panic(err)
	}

	return fw
}

func init() {
	RegisterLoggerWriter(LoggerFile, NewFileWriter)
}
