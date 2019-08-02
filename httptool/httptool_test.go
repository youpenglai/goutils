package httptool

import "testing"

func Test_http(t *testing.T) {
	_, data, err := HttpGet("http://www.baidu.com", map[string]string{})
	if err != nil {
		t.Fatalf("HttpGet err: %v", err)
	}
	t.Logf("len(data): %v", len(data))
}
