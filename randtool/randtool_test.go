package randtool

import "testing"

func Test_int(t *testing.T) {
	t.Log(RandInt64(-100, 50))
}

func Test_strings(t *testing.T) {
	t.Log(Rand32String(12))
	t.Log(Rand64String(12))
}

func Test_uuid(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("uuid: %v", Uuid())
	}
}