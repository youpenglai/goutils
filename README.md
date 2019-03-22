# goutils
Penglai Utils golang version.

Download and install
-------------

    go get -u github.com/youpenglai/goutils


Directory explain
-------------
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 可以使用go test -v测试


Utils directorys
-------------

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| 类型转换 | [convert](convertor/) | 操作数字等 |
| 加密解密 | [crypto](cryptotool/) | base64、md5... |
| HTTP调用 | [httpreq](httptool/http.go) | http调用方式 |
| 随机数据 | [randtool](randtool/) | 获取随机uuid |
| 系统信息 | [system](systool/system.go) | 获取系统信息 |
| token | [token](tokentool/jwt.go) | 安全认证 |