package randtool

import "testing"

func Test_uuid(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("uuid: %v", Uuid())
	}
}
