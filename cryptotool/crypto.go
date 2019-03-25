package cryptotool

import (

	// "fmt"
	"strings"
)

func CreateAuthMap(key, secret, stringToSign string) map[string]string {
	authSignature := HmacSignature(stringToSign, secret)
	authString := strings.Join([]string{key, authSignature}, ":")
	return map[string]string{"auth": authString}
}
