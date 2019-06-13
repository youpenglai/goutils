package tokentool

import (
	"testing"
	"time"
)

func Test_jwt(t *testing.T) {
	pri := NewJwtPri(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALku0jDDorumhCfe
edB145Z7LGWfb5o3TlWRpMwb4hLSlC/mRY480AMpDFYdmmoLtT2GHeMX7koCaBfs
DZ5CkQGMS3k9y1qCZkAH2KBGSLPgCcmK/ZQar2exz6UBKGxwlrAOKlrTcSC+gksG
Xbui1xUPdb2/9mLkZr6grKuu9GZpAgMBAAECgYA+G3HrYPHHNcXFCVyZibBW2+C8
BIkNk9b14+79dayQ24dPuU9CsSqipLW6fBs5zctvYvfLHk+RCtWWkFjQ50pyRBGD
a3IevMnfY7evNqUnT6b44+ivPSKjEIH42BMqMixoxklYVMNBD6Hiblt4Xv8iFau5
A4AQ8b35dAk3ZRCycQJBAOlP3w6qlgM9EwJkxwMySPWBI/O7fuz/goFievEhElH3
16UrFrReZco+NxypDElCx8Efkpm1cHzOpI071EV5ZtMCQQDLMM+/t/dWM/tu8ChL
YDyNSPR6H1aRCXs2POIzS6t+MDYo9fbyikDcSRXR+HA9xawmQAIRapdZ3QO8qcdE
HbBTAkBXEhHCnXMCnmZbvFRxIvELfjh2m4CQ3gJMWv5awubdZEs8PspoYdpHEdTr
g0MsjBmUPz4s0wO58DyE4NElZFs1AkAgVLf5zY/xi3vqfxQqSjnYUU91Tx87HMMa
Mj3b46J6BbnYcqrElPMVGhv6uQlDCv0FaRHrK4bBqcU6c9ldIUw5AkEAm67PxS6r
kIC5oJBbvfoX9XSpxrO5SzCu16tm2J++rA/5SEnIPHe2U89a7SxhSSwMAQ8gTaPq
IvkrNLK6d6FAcA==
-----END PRIVATE KEY-----
`)
	pub := NewJwtPub(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5LtIww6K7poQn3nnQdeOWeyxl
n2+aN05VkaTMG+IS0pQv5kWOPNADKQxWHZpqC7U9hh3jF+5KAmgX7A2eQpEBjEt5
PctagmZAB9igRkiz4AnJiv2UGq9nsc+lAShscJawDipa03EgvoJLBl27otcVD3W9
v/Zi5Ga+oKyrrvRmaQIDAQAB
-----END PUBLIC KEY-----
`)

	token, err := pri.CreateToken("abcdefghj", 2 * time.Second)
	if err != nil {
		t.Fatalf("创建token失败: %v", err.Error())
	}
	t.Log("token:", token)

	uuid, err := pub.VerifyToken(token)
	if err != nil {
		t.Fatalf("验证token失败: %v", err.Error())
	}
	t.Log("uuid:", uuid)

	//time.Sleep(3 * time.Second)
	uuid, err = pub.VerifyToken(token)
	if err != nil {
		t.Fatalf("验证token失败: %v", err.Error())
	}
	t.Log("uuid:", uuid)
}
