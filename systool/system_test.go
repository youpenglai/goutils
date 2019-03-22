package systool

import "testing"

func Test_system_ipMac(t *testing.T) {
	ipList, macList, err := GetIpMac()
	if err != nil {
		t.Fatalf("GetIpMac: %v", err.Error())
	}
	t.Logf("ipList: %v", ipList)
	t.Logf("macList: %v", macList)

	mid, uuid, err := GetMIDAndUUID()
	if err != nil {
		t.Fatalf("GetMIDAndUUID: %v", err.Error())
	}
	t.Logf("mid: %v", mid)
	t.Logf("uuid: %v", uuid)
}

func Test_system_ip(t *testing.T) {
	ip, err := GetExternalIP()
	if err != nil {
		t.Fatalf("GetExternalIP: %v", err.Error())
	}
	t.Logf("ip: %v", ip)
}

func Test_system_uuid(t *testing.T) {
	mid, uuid, err := GetMIDAndUUID()
	if err != nil {
		t.Fatalf("GetMIDAndUUID: %v", err.Error())
	}
	t.Logf("mid: %v", mid)
	t.Logf("uuid: %v", uuid)
}
