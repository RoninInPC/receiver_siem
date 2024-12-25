package hostinfo

import (
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"net"
	"receiver_siem/entity/subject"
	"receiver_siem/hash"
)

var (
	hostName string
	hostOS   string
	ips      []string
)

type HostInfo struct {
	HostName string
	HostOS   string
	IPs      []string
}

func (h HostInfo) JSON() string {
	bytes, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (h HostInfo) Type() subject.SubjectType {
	return subject.HostT
}

func (h HostInfo) Name() string {
	return h.HostName
}

func (h HostInfo) Hash(hash hash.Hash) string {
	return hash(h.JSON())
}

func HostInfoInit() {
	info, _ := host.Info()
	hostName = info.Hostname
	hostOS = info.OS
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
		ips}
}
