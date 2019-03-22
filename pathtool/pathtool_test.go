package pathtool

import "testing"

func Test_pwd(t *testing.T) {
	path, err := GetCurrentPath()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("currentPath: %v", path)

	dirPath, err := GetCurrentDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("currentDir: %v", dirPath)
}

func Test_dir(t *testing.T) {
	b, err := DirExists("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("file exists: %v", b)

	b, err = FileExists("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool/dir.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("file exists: %v", b)

	dirList, err := GetDirFiles("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("dirList: %v", dirList)

	dirList, err = GetDirFilesForSuffixs("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool",
		[]string{".go", ".cpp"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("dirList: %v", dirList)

}
