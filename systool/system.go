package systool

import (
	"github.com/youpenglai/goutils/cmdtool"
	"io/ioutil"
	"net"
	"net/http"
	"sort"
	"strings"
)

// 获取IP列表/MAC列表
func GetIpMac() (ipList []string, macList []string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, inter := range interfaces {
		if inter.HardwareAddr.String() != "" {
			macList = append(macList, inter.HardwareAddr.String())
		}
		var addrs []net.Addr
		addrs, err = inter.Addrs()
		if err != nil {
			return
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ipList = append(ipList, addr.String())
				}
			}
		}
	}
	sort.Strings(ipList)
	sort.Strings(macList)
	return
}

// 获取服务器外网IP
func GetExternalIP() (ip string, err error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

// 获取主板序列号和UUID(需要root权限才能获取到)
func GetMIDAndUUID() (mid string, uuid string, err error) {
	execCmd := `dmidecode |grep -A 10 "System Information" | grep "Serial Number\|UUID"`
	outStr, err := cmdtool.Cmd(execCmd)
	if err != nil {
		return
	}

	getSubField := func(str string) string {
		if n := strings.Index(outStr, str); n >= 0 {
			n += len(str)
			m := strings.IndexByte(outStr[n:], '\n')
			if m != -1 {
				return outStr[n:][:m]
			} else {
				return outStr[n:]
			}
		}
		return ""
	}

	return getSubField("Serial Number: "), getSubField("UUID: "), nil
}
