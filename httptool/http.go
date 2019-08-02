package httptool

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

func HttpGet(url string, headers map[string]string) (httpCode int, result []byte, err error) {
	return httpCall("GET", url, headers, nil)
}

func HttpPost(url string, headers map[string]string, data []byte) (httpCode int, result []byte, err error) {
	return httpCall("POST", url, headers, bytes.NewBuffer(data))
}

func HttpPut(url string, headers map[string]string, data []byte) (httpCode int, result []byte, err error) {
	return httpCall("PUT", url, headers, bytes.NewBuffer(data))
}

func HttpDelete(url string, headers map[string]string, data []byte) (httpCode int, result []byte, err error) {
	return httpCall("DELETE", url, headers, bytes.NewBuffer(data))
}

func httpCall(method string, url string, headers map[string]string, body io.Reader) (httpCode int, result []byte, err error) {
	if _, ok := methods[method]; !ok {
		err = errors.New("Http Call Method Must In[GET, POST, PUT, DELETE]")
		return
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

	httpCode = resp.StatusCode
	result, err = ioutil.ReadAll(resp.Body)
	return
}
