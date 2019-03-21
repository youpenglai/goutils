package httpreq

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var methods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
}

func HttpGet(url string, headers map[string]string) (result []byte, err error) {
	return httpCall("GET", url, headers, nil)
}

func HttpPost(url string, headers map[string]string, data []byte) (result []byte, err error) {
	return httpCall("POST", url, headers, bytes.NewBuffer(data))
}

func HttpPut(url string, headers map[string]string, data []byte) (result []byte, err error) {
	return httpCall("PUT", url, headers, bytes.NewBuffer(data))
}

func HttpDelete(url string, headers map[string]string, data []byte) (result []byte, err error) {
	return httpCall("DELETE", url, headers, bytes.NewBuffer(data))
}

func httpCall(method string, url string, headers map[string]string, body io.Reader) (result []byte, err error) {
	if _, ok := methods[method]; !ok {
		return nil, errors.New("Http Call Method Must In[GET, POST, PUT, DELETE]")
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	for name, value := range headers {
		req.Header.Set(name, value)
	}

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
