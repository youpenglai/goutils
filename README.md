# goutils
Penglai Utils golang version.

Download and install
-------------

    go get -u github.com/youpenglai/goutils

convert:
-------------
```go
    convert.ConvertInt64("16")
    convert.TimeToStr(time.Now())
    ...
```

crypto:
-------------
```go
    base64、aes、des、md5、rsa、sha、hmac、crypto
```

httpreq:
-------------
```go
    HttpGet(url, headers)
    HttpPost(url, headers, data)
    ...
```

token:
-------------
```go
    pri := NewJwtPri(priKey)
    pri.CreateToken()

    pub := NewJwtPub(pubKey)
    pub.VerifyToken()
```

uuid:
-------------
```go
    Uuid()
```

system:
-------------
```go
    ipList, ipMac, err := GetIpMac()
    ip, err := GetExternalIP()
    mid, uuid ,err := GetMIDAndUUID()
```