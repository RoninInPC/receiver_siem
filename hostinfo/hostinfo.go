package hostinfo

import (
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"net"
	"receiver_siem/hash"
)

var (
	hostName string
	hostOS   string
	codeName string
	ips      []string
)

type HostInfo struct {
	HostName string
	HostOS   string
	CodeName string
	IPs      []string
}

func (h HostInfo) JSON() string {
	bytes, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (h HostInfo) Name() string {
	return h.HostName
}

func (h HostInfo) Hash(hash hash.Hash) string {
	return hash(h.JSON())
}

func HostInfoInit(key string) {
	info, _ := host.Info()
	hostName = info.Hostname
	hostOS = info.OS
	codeName = key
	ips, _ = getLocalIPs()
}

func getLocalIPs() ([]string, error) {
	var ips []string
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, nil
}

func GetHostInfo() HostInfo {
	return HostInfo{
		hostName,
		hostOS,
		codeName,
		ips}
}
