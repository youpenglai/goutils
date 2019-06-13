package cryptotool

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func GetSha1(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetSha256(data string) string {
	h := sha256.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetSha512(data string) string {
	h := sha512.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}