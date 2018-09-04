package iputils

import (
	"net"
)

// GetPublicIPv4s get list of public ipv4s
func GetPublicIPv4s() []net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	ips := []net.IP{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && IsPublicIPv4(ip) {
				ips = append(ips, ip)
			}
		}
	}
	return ips
}

// IsPublicIPv4 check ip is public ipv4
func IsPublicIPv4(ip net.IP) bool {
	ip4 := ip.To4()
	if ip4 == nil ||
		ip.IsLoopback() ||
		ip.IsLinkLocalMulticast() ||
		ip.IsLinkLocalUnicast() ||
		ip4[0] == 10 ||
		ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 ||
		ip4[0] == 192 && ip4[1] == 168 {
		return false
	}
	return true
}
