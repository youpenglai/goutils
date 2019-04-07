package logger

import (
	"github.com/youpenglai/goutils/pathtool"
	"os"
	"sync"
)

type FileWriter struct {
	mu       sync.Mutex
	filename string

	rotate   bool
	maxSize  int64
	maxLines int
	file     *os.File
	wb       int64
}

func (fw *FileWriter) openLogFile() {
	pathtool.FileExists(fw.filename)
}

func (fw *FileWriter) Colorful() bool {
	return false
}

func (fw *FileWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (fw *FileWriter) Flush() {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.file.Sync()
}
