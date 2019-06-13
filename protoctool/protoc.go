package protoctool

import (
	"fmt"
	"github.com/youpenglai/goutils/cmdtool"
)

// 生成proto.pb.go文件（相对路径）
func Protoc(srcDir string) error {
	_, err := cmdtool.Cmd(fmt.Sprintf(`protoc --go_out=. %s/*.proto`, srcDir))
	return err
}

// 生成proto.pb.go文件（相对路径）（包括grpc函数）
func ProtocGRPC(srcDir string) error {
	_, err := cmdtool.Cmd(fmt.Sprintf(`protoc --go_out=plugins=grpc:. %s/*.proto`, srcDir))
	return err
}
