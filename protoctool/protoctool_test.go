package protoctool

import (
	"github.com/gpmgo/gopm/modules/log"
	"testing"
)

func Test_protoc(t *testing.T) {
	if err := Protoc("", ""); err != nil {
		log.Fatal("err: %v", err.Error())
	}
}