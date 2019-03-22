package protoctool

import (
	"fmt"
	"github.com/youpenglai/goutils/cmdtool"
	"github.com/youpenglai/goutils/pathtool"
	"strings"
)

// 生成proto.pb.go文件
func Protoc(srcDir string, destDir string) error {
	files, err := pathtool.GetDirFilesForSuffixs(srcDir, []string{".proto"})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf(`protoc --go_out=%s %s`, destDir, strings.Join(files, " ")))
	_, err = cmdtool.Cmd(fmt.Sprintf(`protoc --go_out=%s %s`, destDir, strings.Join(files, " ")))
	return err
}

// 生成proto.pb.go文件（包括grpc函数）
func ProtocGRPC(srcDir string, destDir string) error {
	files, err := pathtool.GetDirFiles(srcDir)
	if err != nil {
		return err
	}

	_, err = cmdtool.Cmd(fmt.Sprintf("protoc --go_out=plugins=grpc:%s %s", destDir, strings.Join(files, " ")))
	return err
}
