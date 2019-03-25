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
| 命令行 | [cmd](comtool/) | 调用系统命令行 |
| 类型转换 | [convert](convertor/) | 数字/时间类型转换 |
| 加密解密 | [crypto](cryptotool/) | base64、md5...多种加密解密 |
| HTTP调用 | [http](httptool/http.go) | http快速调用 |
| 路径文件 | [path](pathtool/) | 当前路径、文件列表 |
| protobuf | [protobuf](protoctool/) | 生成protobuf的go文件 |
| 随机数据 | [rand](randtool/) | 获取随机数、uuid |
| 系统信息 | [system](systool/system.go) | 获取系统信息 |
| token | [token](tokentool/jwt.go) | 安全认证 |