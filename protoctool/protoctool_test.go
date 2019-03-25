package protoctool

import (
	"testing"
)

func Test_protoc(t *testing.T) {

	if err := Protoc("helloworld/"); err != nil {
		t.Fatal(err.Error())
	}

	if err := ProtocGRPC("helloworld/"); err != nil {
		t.Fatal(err.Error())
	}
}
