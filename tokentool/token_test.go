package tokentool

import (
	"testing"
	"time"
)

func Test_jwt(t *testing.T) {
	pri := NewJwtPri(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCCDhEPhxbfD0hXnP4r4W6tuSFAEFXnL0KXwqV3/EMA0dwnjmC+
q7861sGZq2frc1zPN4kMjFlTqzcYqZIhm0oEcA4FVFTBMANWZ5urnm1GcxdPPLaT
CTAIBFJd3wOuz1yyVw8CIviHXVDr5105QtLLMQCf/GRytGvStb7xm283MQIDAQAB
AoGAQ7GOr8Zg4TnTLcdJhQ2RzlrOM8pM7EhtIRDvj8R+G68gKHazZIuAvd+KZbIS
JTum0zTU4PMsi6BSzUN2DIgHvgSI3Ea42dgXg8Kk2E+E9JNdxsacF0R8TkDyieRQ
EGcnK7RlFafgPfE2usW/PfLVNxAUgWaJz3wSDYVa6t/QDQECQQDl66PgCCkqWnDL
9EBLo11isjQHidWqmP7neNLROkmTlMNW/wEYeSEOaBhbYpBmsq5O7uUNP7s6FORs
hnLAKF5fAkEAkM6P7S0hQsSLmHeqz3J3DOUmdjvilqNh/STedGBLXPwlETox7zJE
jmTk3THdzCG4OG+85H4gnpc4E5qYWh00bwJBAJQxMlnWyLmsu2Ep2DsoW6uekpx9
QcOIg1usa61BtWSzEMjE6e7dO+ouO+zC9bHL3z+vCOFmP4XrK2OrsNCb9F8CQGJD
2kVvn6eIatZ/Nmlp2sHkarJHNx9UJfW75D2C7wFdleXX4PTZ3s3mR6yW31Vb+IBL
yxofQRF9uq8K9KIKOX8CQEbKfxmMyCOtL3ugF3apyNE0cspdUhvjUkRooGdIJYN7
d4BVle4r2sVL4Vq2MgXCYu10eduxuPgiRXd1F8Gn8sQ=
-----END RSA PRIVATE KEY-----`)
	pub := NewJwtPub(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCDhEPhxbfD0hXnP4r4W6tuSFA
EFXnL0KXwqV3/EMA0dwnjmC+q7861sGZq2frc1zPN4kMjFlTqzcYqZIhm0oEcA4F
VFTBMANWZ5urnm1GcxdPPLaTCTAIBFJd3wOuz1yyVw8CIviHXVDr5105QtLLMQCf
/GRytGvStb7xm283MQIDAQAB
-----END PUBLIC KEY-----`)

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
