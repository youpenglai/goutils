package pathtool

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取当前执行文件的绝对路径
func GetCurrentPath() (path string, err error) {
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	path = strings.Replace(path, "\\", "/", -1) // 将\替换成/
	return
}

// 获取当前执行文件的上一级目录地址
func GetCurrentDir() (dir string, err error) {
	path, err := GetCurrentPath()
	if err != nil {
		return
	}

	return substr(path, 0, strings.LastIndex(path, "/")), nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}