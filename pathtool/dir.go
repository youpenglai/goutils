package pathtool

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// 文件夹是否存在
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, errors.New("Path is not Dir")
	}
	return true, nil
}

// 文件是否存在
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, errors.New("Path is Dir")
	}
	return true, nil
}

// 获取文件夹下所有文件（递归、无文件夹）
func GetDirFiles(dirPath string) (dirList []string, err error) {
	err = filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				dirList = append(dirList, path)
			}
			return nil
		},
	)
	return
}

// 获取文件夹下符合后缀名的所有文件（递归、无文件夹）
func GetDirFilesForSuffixs(dirPath string, suffixs []string) (dirList []string, err error) {
	err = filepath.Walk(dirPath,
		func(path string, f os.FileInfo, err error) error {
			var flag = false
			for _, suffix := range suffixs {
				flag = flag || strings.HasSuffix(path, suffix)
			}

			if !f.IsDir() && (flag) {
				dirList = append(dirList, path)
			}

			return nil
		},
	)
	return
}

/*
// 获取文件夹下所有文件（递归、无文件夹）
func GetDirFiles(dirPath string) (dirList []string, err error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(filepath.Join(dirPath, file.Name()))
			dirSubList, e := WalkGetDirFiles(filepath.Join(dirPath, file.Name()))
			if e != nil {
				err = e
				return
			}
			dirList = append(dirList, dirSubList...)
		} else {
			dirList = append(dirList, filepath.Join(dirPath, file.Name()))
		}
	}
	return
}
*/

/*
// 获取文件夹下符合后缀名的所有文件（递归、无文件夹）
func WalkGetDirFilesForSuffixs(dirPath string, suffixs []string) (dirList []string, err error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			dirSubList, e := WalkGetDirFilesForSuffixs(filepath.Join(dirPath, file.Name()), suffixs)
			if e != nil {
				err = e
				return
			}
			dirList = append(dirList, dirSubList...)

		} else {
			var flag = false
			for _, suffix := range suffixs {
				flag = flag || strings.HasSuffix(dirPath, suffix)
			}
			if flag {
				dirList = append(dirList, filepath.Join(dirPath, file.Name()))
			}
		}
	}
	return
}
*/