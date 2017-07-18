package iputils

import (
	"fmt"
	"net"
)

// ContainsIP checks if 'ip' is contained inside one of the 'networks' provided.
// It returns 'ok' when the the 'ip' is contained. Also it returns the 'network'
// where it's contained.
// Returns error if the ip or the networks can't be parsed.
func ContainsIP(ip string, networks ...string) (ok bool, network string, err error) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false, "", fmt.Errorf("%v is not a valid IP address", ip)
	}

	for _, network = range networks {
		_, parsedNet, err := net.ParseCIDR(network)
		if err != nil {
			return false, "", fmt.Errorf("error parsing the network: %v", err)
		}

		if parsedNet.Contains(parsedIP) {
			return true, network, nil
		}
	}

	return false, "", nil
}

// ExpandCIDR returns a list of the IPs contained in the given CIDR range.
// If the removeNetAndBroadcast is true, the network and broadcast IP addresses
// are not returned.
func ExpandCIDR(network string, removeNetAndBroadcast bool) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	if removeNetAndBroadcast {
		ips = ips[1 : len(ips)-1]
	}

	return ips, nil
}

// From https://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
