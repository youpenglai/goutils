# goutils
Penglai Utils golang version.

Download and install
-------------

    go get -u github.com/youpenglai/goutils


directory explain
-------------
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 可以使用go test -v测试


utils directorys
-------------

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| 类型转换 | [convert](convert/) | 操作数字等 |
| 加密解密 | [crypto](crypto/) | base64、md5... |
| HTTP调用 | [httpreq](httpreq/http.go) | http调用方式 |
| token | [token](token/jwt.go) | 安全认证 |
| uuid | [uuid](uuid.go) |  获取随机uuid |
| system | [system](system.go) | 获取信息信息 |