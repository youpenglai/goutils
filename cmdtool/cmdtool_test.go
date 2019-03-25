package cmdtool

import "testing"

func Test_cmd(t *testing.T) {
	outStr, err := Cmd("ls")
	if err !=nil {
		t.Fatal(err.Error())
	}
	t.Logf("ls: [%s]", outStr)
}