package protoctool

import (
	"testing"
)

func Test_protoc(t *testing.T) {

	if err := Protoc("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool",
		"/home/embiid/go/src/github.com/youpenglai/goutils/pathtool"); err != nil {
		t.Fatal(err.Error())
	}
	/*
		if err := ProtocGRPC("/home/embiid/go/src/github.com/youpenglai/goutils/pathtool",
			"/home/embiid/go/src/github.com/youpenglai/goutils/pathtool"); err != nil {
			t.Fatal(err.Error())
		}
	*/
}
